package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics interface {
	CreateSuccessInc()
	UpdateSuccessInc()
	RemoveSuccessInc()
}

type metrics struct {
	createCounter prometheus.Counter
	updateCounter prometheus.Counter
	removeCounter prometheus.Counter
}

func NewMetrics() Metrics {
	m := &metrics{
		createCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "create_course_response_success",
			Help: "Create course response success",
		}),
		updateCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "update_course_response_success",
			Help: "Update course response success",
		}),
		removeCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "remove_course_response_success",
			Help: "Remove course response success",
		}),
	}

	prometheus.MustRegister(m.createCounter)
	prometheus.MustRegister(m.updateCounter)
	prometheus.MustRegister(m.removeCounter)

	return m
}

func (m *metrics) CreateSuccessInc() {
	m.createCounter.Inc()
}

func (m *metrics) UpdateSuccessInc() {
	m.updateCounter.Inc()
}

func (m *metrics) RemoveSuccessInc() {
	m.removeCounter.Inc()
}
