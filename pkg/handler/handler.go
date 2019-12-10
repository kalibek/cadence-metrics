package handler

import (
	"epam.com/cadence-metrics/pkg/collector"
	"epam.com/cadence-metrics/pkg/config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func MetricsHandler(c *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metrics, err := collector.CollectMetrics(c)
		if err != nil {
			log.Errorf("cannot collect cadence metrics %v", err)
		}
		for _, m := range metrics {
			m.WriteMetric(w)
		}
	}
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("requested %s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}
