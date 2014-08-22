package main

import (
	"github.com/rcrowley/go-metrics"
)

type Metrics struct {
}

func NewMetrics() (*Metrics, error) {
	metrics := &Metrics{}
	return metrics, nil
}

func (m *Metrics) TrackMeter(key string, value int64) error {
	counter := metrics.GetOrRegisterMeter(key, metrics.DefaultRegistry)
	counter.Mark(value)
	return nil
}

func (m *Metrics) TrackGauge(key string, value int64) error {
	counter := metrics.GetOrRegisterGauge(key, metrics.DefaultRegistry)
	counter.Update(value)
	return nil
}

func (m *Metrics) TrackCounter(key string, value int64) error {
	counter := metrics.GetOrRegisterCounter(key, metrics.DefaultRegistry)
	counter.Inc(value)
	return nil
}
