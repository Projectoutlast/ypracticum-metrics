package restapi

import (
	v1 "github.com/Projectoutlast/ypracticum-metrics/internal/handler/restapi/v1"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
	"github.com/go-chi/chi/v5"
)

func NewRouter(mux *chi.Mux, m v1.Metric, l logger.Interface) {
	v1.NewRoutes(mux, m, l)
}
