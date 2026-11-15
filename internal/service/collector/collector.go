package collector

import (
	"context"
	"math/rand/v2"
	"runtime"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
	"github.com/Projectoutlast/ypracticum-metrics/internal/repository"
	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
)

type UseCase struct {
	repo repository.Collector
}

func New(repo repository.Collector) service.Collector {
	return &UseCase{repo: repo}
}

func (uc *UseCase) Collect(ctx context.Context) error {
	newMetrics := collectMetrics()

	uc.repo.Collect(ctx, newMetrics)

	return nil
}

func collectMetrics() map[string]*models.Metrics {
	var stats runtime.MemStats

	runtime.ReadMemStats(&stats)

	return map[string]*models.Metrics{
		"Alloc":         metric("Alloc", "gauge", float64(stats.Alloc)),
		"BuckHashSys":   metric("BuckHashSys", "gauge", float64(stats.BuckHashSys)),
		"Frees":         metric("Frees", "gauge", float64(stats.Frees)),
		"GCCPUFraction": metric("GCCPUFraction", "gauge", stats.GCCPUFraction),
		"GCSys":         metric("GCSys", "gauge", float64(stats.GCSys)),
		"HeapAlloc":     metric("HeapAlloc", "gauge", float64(stats.HeapAlloc)),
		"HeapIdle":      metric("HeapIdle", "gauge", float64(stats.HeapIdle)),
		"HeapInuse":     metric("HeapInuse", "gauge", float64(stats.HeapInuse)),
		"HeapObjects":   metric("HeapObjects", "gauge", float64(stats.HeapObjects)),
		"HeapReleased":  metric("HeapReleased", "gauge", float64(stats.HeapReleased)),
		"HeapSys":       metric("HeapSys", "gauge", float64(stats.HeapSys)),
		"LastGC":        metric("LastGC", "gauge", float64(stats.LastGC)),
		"Lookups":       metric("Lookups", "gauge", float64(stats.Lookups)),
		"MCacheInuse":   metric("MCacheInuse", "gauge", float64(stats.MCacheInuse)),
		"MCacheSys":     metric("MCacheSys", "gauge", float64(stats.MCacheInuse)),
		"MSpanInuse":    metric("MSpanInuse", "gauge", float64(stats.MSpanInuse)),
		"MSpanSys":      metric("MSpanSys", "gauge", float64(stats.MSpanSys)),
		"Mallocs":       metric("Mallocs", "gauge", float64(stats.Mallocs)),
		"NextGC":        metric("NextGC", "gauge", float64(stats.NextGC)),
		"NumForcedGC":   metric("NumForcedGC", "gauge", float64(stats.NumForcedGC)),
		"NumGC":         metric("NumGC", "gauge", float64(stats.NumGC)),
		"OtherSys":      metric("OtherSys", "gauge", float64(stats.OtherSys)),
		"PauseTotalNs":  metric("PauseTotalNs", "gauge", float64(stats.PauseTotalNs)),
		"StackInuse":    metric("StackInuse", "gauge", float64(stats.StackInuse)),
		"StackSys":      metric("StackSys", "gauge", float64(stats.StackSys)),
		"Sys":           metric("Sys", "gauge", float64(stats.Sys)),
		"TotalAlloc":    metric("TotalAlloc", "gauge", float64(stats.TotalAlloc)),
		"RandomValue":   metric("RandomValue", "gauge", rand.Float64()),
		"PollCount":     metric("PollCount", "counter", float64(1)),
	}
}

func metric(id string, mtype string, v float64) *models.Metrics {
	return &models.Metrics{ID: id, MType: mtype, Value: new(v)}
}
