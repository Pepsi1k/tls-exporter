package main

import (
	"log"
	"net/http"

	// "github.com/aws/aws-sdk-go-v2/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// "crypto/tls"
)

var (
  domainsList = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
      Name: "domain",
      Help: "just domain",
    }, []string{"domain"},
  )
)

func main() {
  prometheus.MustRegister(domainsList)

  domains := [...]string{
    "cppstudioy.com",
    "elegro.eu",
    "dashboard-stage.niko.technology",
    "expired-ecc-dv.ssl.com",
  }

  // for i := 0; i < len(domains); i++ {
  //   fmt.Println(domains[i])
  // }


  for i := 0; i < len(domains); i++ {
    domainsList.With(prometheus.Labels{"domain": domains[i]}).Set(0)
  }


  http.Handle("/metrics", promhttp.Handler())
  
  if err := http.ListenAndServe(":80", nil); err != nil {
    log.Fatal(err)
  }
  
}