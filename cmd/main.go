package main

import (
	"crypto/tls"
	"log"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/spf13/viper"
	bulkdatacollectorservices "github.com/zdrgeo/bulk-data-collector/pkg/services"
	"github.com/zdrgeo/cwmp-interceptor/pkg/handlers"
	"github.com/zdrgeo/cwmp-interceptor/pkg/services"
)

var (
	logger *slog.Logger
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
}

func main() {
	targetURL, err := url.Parse(viper.GetString("TARGET_URL"))

	if err != nil {
		log.Fatal(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	reverseProxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	collectorService := bulkdatacollectorservices.NewMockCollectorService()
	eavesdropperService := services.NewEavesdropperService(collectorService, nil)
	interceptorHandler := handlers.NewInterceptorHandler(targetURL, reverseProxy, eavesdropperService)

	http.Handle("/", http.HandlerFunc(interceptorHandler.Intercept))

	if err := http.ListenAndServe(":8880", nil); err != nil {
		log.Fatal(err)
	}
}
