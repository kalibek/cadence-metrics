package collector

import (
	"epam.com/cadence-metrics/pkg/config"
	"epam.com/cadence-metrics/pkg/model"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strconv"
	"strings"
)

const (
	CommandFmt = "/usr/local/bin/cadence --ad %s:%s --do %s tasklist desc --tl %s | grep @ | wc -l"
)

func newCadencePoller(value float32, app string) *model.Metric {
	return &model.Metric{
		Name:  "cadence_poller_count",
		Help:  "cadence poller count",
		Type:  "gauge",
		Value: value,
		Tags: []model.Tag{
			{Name: "application", Value: app},
		},
	}
}

func CollectMetrics(c *config.Config) ([]model.MetricWriter, error) {
	metrics := cadencePollers(c)
	return metrics, nil
}

func cadencePollers(c *config.Config) []model.MetricWriter {
	metrics := make([]model.MetricWriter, 0)
	for _, tl := range c.Cadence.TaskList {
		command := fmt.Sprintf(CommandFmt, c.Cadence.Server, c.Cadence.Port, c.Cadence.Domain, tl)
		out, err := exec.Command("/bin/sh", "-c", command).Output()

		if err == nil {
			val, err := strconv.ParseInt(strings.TrimSpace(string(out)), 10, 32)
			if err != nil {
				log.Errorf("cannot get metrics %s for value %v", tl, string(out))
			}
			m := newCadencePoller(float32(val), c.Cadence.Server)
			metrics = append(metrics, model.MetricWriter(m))
		} else {
			log.Errorf("error while counting pollers %v", err)
		}

	}
	return metrics
}
