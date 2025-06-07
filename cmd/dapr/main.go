package main

import (
	"crypto/tls"
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"

	daprclient "github.com/dapr/go-sdk/client"
	"github.com/spf13/viper"
	daprbulkdatacollectorservices "github.com/zdrgeo/bulk-data-collector/pkg/services/dapr"
	"github.com/zdrgeo/cwmp-interceptor/pkg/handlers"
	"github.com/zdrgeo/cwmp-interceptor/pkg/services"
)

const (
	storeName  = "iotoperations-statestore"
	pubSubName = "iotoperations-pubsub"
	topicName  = "collector"
)

var (
	logger     *slog.Logger
	daprClient daprclient.Client
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

	initDapr()
}

func initDapr() {
	var err error

	daprClient, err = daprclient.NewClient()

	if err != nil {
		log.Panic(err)
	}
}

func main() {
	mainDapr()
}

func mainDapr() {
	targetURL, err := url.Parse(viper.GetString("TARGET_URL"))

	if err != nil {
		log.Panic(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	reverseProxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	collectorServiceOptions := &daprbulkdatacollectorservices.DaprCollectorServiceOptions{
		PubSubName: pubSubName,
		TopicName:  topicName,
	}

	collectorService := daprbulkdatacollectorservices.NewDaprCollectorService(daprClient, collectorServiceOptions)

	eavesdropperService := services.NewEavesdropperService(collectorService, nil)
	eavesdropperHandler := handlers.NewEavesdropperHandler(eavesdropperService)
	interceptorHandler := handlers.NewInterceptorHandler(targetURL, reverseProxy, eavesdropperService)

	http.Handle("/eavesdropper", http.HandlerFunc(eavesdropperHandler.Eavesdrop))
	http.Handle("/interceptor", http.HandlerFunc(interceptorHandler.Intercept))

	if err := http.ListenAndServe(":8880", nil); err != nil && err != http.ErrServerClosed {
		log.Panic(err)
	}
}
