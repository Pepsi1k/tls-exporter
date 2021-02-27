package main

import (
	"log"
	"net/http"

  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func recordMetrics() {
  go func() {
    for {
      opsProcessed.Inc()
      time.Sleep
    }
  }
}

func main() {

	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
