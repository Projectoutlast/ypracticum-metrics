package restapi

import (
	v1 "github.com/Projectoutlast/ypracticum-metrics/internal/handler/restapi/v1"
	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
	"github.com/go-chi/chi/v5"
)

func NewRouter(mux *chi.Mux, m service.Metric, l logger.Interface) {
	v1.NewRoutes(mux, m, l)
}
