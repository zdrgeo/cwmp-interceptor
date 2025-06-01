package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	azureeventhubsbulkdatacollectorservices "github.com/zdrgeo/bulk-data-collector/pkg/services/azureeventhubs"
	"github.com/zdrgeo/cwmp-interceptor/pkg/handlers"
	"github.com/zdrgeo/cwmp-interceptor/pkg/services"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var (
	logger     *slog.Logger
	credential *azidentity.DefaultAzureCredential
)

func init() {
	logger = slog.Default()
	// Use otelslog bridge to integrate with OpenTelemetry (https://pkg.go.dev/go.opentelemetry.io/otel/sdk/log)
	// logger := slog.New(slog.NewTextHandler(nil, &slog.HandlerOptions{AddSource: true}))
	// logger := slog.New(slog.NewJSONHandler(nil, &slog.HandlerOptions{AddSource: true}))

	viper.AddConfigPath(".")
	// viper.SetConfigFile(".env")
	// viper.SetConfigName("config")
	// viper.SetConfigType("env") // "env", "json", "yaml"
	viper.SetEnvPrefix("iridium")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	initAzureEventHubs()
}

func initAzureEventHubs() {
	var err error

	encoder := json.NewEncoder(os.Stdout)

	encoder.SetIndent("", "  ")

	stdoutExporter, err := stdoutmetric.New(
		stdoutmetric.WithEncoder(encoder),
		stdoutmetric.WithoutTimestamps(),
	)

	if err != nil {
		log.Panic(err)
	}

	_ = stdoutExporter

	otlpgrpcExporter, err := otlpmetricgrpc.New(
		context.Background(),
		otlpmetricgrpc.WithEndpoint("localhost:4317"),
		otlpmetricgrpc.WithInsecure(),
	)

	if err != nil {
		log.Panic(err)
	}

	_ = otlpgrpcExporter

	otlphttpExporter, err := otlpmetrichttp.New(
		context.Background(),
		otlpmetrichttp.WithEndpoint("localhost:4318"),
		otlpmetrichttp.WithInsecure(),
	)

	if err != nil {
		log.Panic(err)
	}

	_ = otlphttpExporter

	resource := resource.NewSchemaless(
		semconv.ServiceName("iridium-collector"),
	)

	_ = resource

	periodicReader := metric.NewPeriodicReader(
		stdoutExporter,
		// otlpgrpcExporter,
		// otlphttpExporter,
		metric.WithInterval(10*time.Second),
	)

	_ = periodicReader

	prometheusExporter, err := prometheus.New()

	if err != nil {
		log.Fatal(err)
	}

	_ = prometheusExporter

	meterProvider := metric.NewMeterProvider(
		// metric.WithResource(resource),
		// metric.WithReader(periodicReader),
		metric.WithReader(prometheusExporter),
	)

	otel.SetMeterProvider(meterProvider)

	credential, err = azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		log.Panic(err)
	}
}

func main() {
	mainAzureEventHubs()
}

func mainAzureEventHubs() {
	targetURL, err := url.Parse(viper.GetString("TARGET_URL"))

	if err != nil {
		log.Fatal(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	reverseProxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	ctx := context.Background()

	// producerClient, err := azeventhubs.NewProducerClient(viper.GetString("AZURE_EVENTHUBS_NAMESPACE"), viper.GetString("AZURE_EVENTHUBS_EVENTHUB"), credential, nil)
	producerClient, err := azeventhubs.NewProducerClientFromConnectionString(viper.GetString("AZURE_EVENTHUBS_CONNECTION_STRING"), viper.GetString("AZURE_EVENTHUBS_EVENTHUB"), nil)

	if err != nil {
		log.Panic(err)
	}

	defer producerClient.Close(ctx)

	collectorServiceOptions := &azureeventhubsbulkdatacollectorservices.AzureEventHubsCollectorServiceOptions{
		PartitionQueueLimit:     viper.GetInt("PARTITION_QUEUE_LIMIT"),
		PartitionProducersCount: viper.GetInt("PARTITION_PRODUCERS_COUNT"),
	}

	collectorService, err := azureeventhubsbulkdatacollectorservices.NewAzureEventHubsCollectorService(producerClient, collectorServiceOptions)

	if err != nil {
		log.Panic(err)
	}

	eavesdropperService := services.NewEavesdropperService(collectorService, nil)
	eavesdropperHandler := handlers.NewEavesdropperHandler(eavesdropperService)
	interceptorHandler := handlers.NewInterceptorHandler(targetURL, reverseProxy, eavesdropperService)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/eavesdropper", http.HandlerFunc(eavesdropperHandler.Eavesdrop))
	http.Handle("/interceptor", http.HandlerFunc(interceptorHandler.Intercept))

	runErr := make(chan error)

	go func() {
		runErr <- collectorService.Run(ctx)
	}()

	listenAndServeErr := http.ListenAndServe(":8088", nil)

	if err := <-runErr; err != nil {
		log.Panic(err)
	}

	if listenAndServeErr != nil {
		log.Panic(err)
	}
}
