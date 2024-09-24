package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func main() {
	// Create a Prometheus exporter
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}

	// Create a meter provider with the Prometheus exporter
	provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	defer provider.Shutdown(context.Background())

	// Set the global meter provider
	otel.SetMeterProvider(provider)

	// Get a meter
	meter := provider.Meter("my-service")

	// Create a counter
	counter, err := meter.Int64Counter("my_counter", metric.WithDescription("A sample counter"))
	if err != nil {
		log.Fatal(err)
	}

	// Set up Gin server
	r := gin.Default()

	// Add a route that increments the counter
	r.GET("/increment", func(c *gin.Context) {
		counter.Add(c.Request.Context(), 1)
		c.JSON(http.StatusOK, gin.H{"message": "Counter incremented"})
	})

	// Expose Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
