package v1

import (
	"net/http"

	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
)

func NewRoutes(mux *http.ServeMux, c service.Counter, g service.Gauge, l logger.Interface) {
	h := &V1{c: c, g: g, l: l}

	mux.HandleFunc("/ping", h.ping)
	mux.HandleFunc("/update/{type}/{name}/{value}", h.updateMetrics)
}
