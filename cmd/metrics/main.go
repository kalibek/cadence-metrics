package main

import (
	"epam.com/cadence-metrics/pkg/config"
	"epam.com/cadence-metrics/pkg/handler"
	"flag"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	var port, configPath string

	flag.StringVar(&port, "port", ":9010", "server port")
	flag.StringVar(&configPath, "config", "config.yaml", "config file")

	flag.Parse()

	c, err := config.LoadConfig(configPath)
	if err != nil {
		log.Errorf("could not load config file %v", err)
	}

	r := http.NewServeMux()
	r.HandleFunc("/metrics", handler.LoggingMiddleware(handler.MetricsHandler(c)))
	log.Infof("Started on port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
