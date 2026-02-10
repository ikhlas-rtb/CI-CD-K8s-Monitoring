package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Total HTTP requests",
	},
	[]string{"method", "route", "status_code"},
)

func main() {
	prometheus.MustRegister(httpRequestsTotal)

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		httpRequestsTotal.WithLabelValues("GET", "/health", "200").Inc()
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"service": "go-api",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
})

	r.GET("/api/hello", func(c *gin.Context) {
		httpRequestsTotal.WithLabelValues("GET", "/api/hello", "200").Inc()
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from GO API! 🫡",
			"version": "1.0.0",
			"language": "GO",
	})
})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	_ = r.Run(":8080")
}
