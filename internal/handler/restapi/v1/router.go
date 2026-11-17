package v1

import (
	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
	"github.com/go-chi/chi/v5"
)

func NewRoutes(mux *chi.Mux, m service.Metric, l logger.Interface) {
	h := &V1{m: m, l: l}

	mux.Get("/ping", h.ping)
	mux.Get("/", h.metrics)
	mux.Post("/update/{type}/{name}/{value}", h.updateMetrics)
	mux.Get("/value/{type}/{name}", h.getMetric)
}
