package model

import (
	"bytes"
	"fmt"
	"io"
)

type MetricWriter interface {
	WriteMetric(w io.Writer)
}

type Tag struct {
	Name  string
	Value string
}

type Metric struct {
	Name  string
	Help  string
	Type  string
	Value float32
	Tags  []Tag
}

func (m *Metric) WriteMetric(w io.Writer) {
	help := fmt.Sprintf("# HELP %s %s\n", m.Name, m.Help)
	metricType := fmt.Sprintf("# TYPE %s %s\n", m.Name, m.Type)
	buf := &bytes.Buffer{}
	if len(m.Tags) > 0 {
		buf.WriteString("{")
		for i, t := range m.Tags {
			if i > 0 {
				buf.WriteString(",")
			}
			buf.WriteString(fmt.Sprintf("%s=\"%s\"", t.Name, t.Value))
		}
		buf.WriteString("}")
	}
	value := fmt.Sprintf("%s%s %.2f\n", m.Name, buf.String(), m.Value)
	w.Write([]byte(help))
	w.Write([]byte(metricType))
	w.Write([]byte(value))
}
