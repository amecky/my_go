package view

import (
	"system-info-server/domain"
)

type ListView struct {
	Environment string
	Infos []domain.InfoData
}

type RequestsView struct {
	Environment string
	Service string
	Requests []domain.CustomLogData
}

type MetricsView struct {
	Environment string
	Service string
	Metrics []domain.Metric
}

type PerformanceData struct {
	Value int
	Key string
	RequestTimestamp string
}
type ChartsView struct {
	Path string
	Environment string
	Data []PerformanceData
}