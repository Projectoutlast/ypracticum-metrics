package restapi

import (
	"net/http"

	v1 "github.com/Projectoutlast/ypracticum-metrics/internal/handler/restapi/v1"
	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
)

func NewRouter(mux *http.ServeMux, c service.Counter, g service.Gauge, l logger.Interface) {
	v1.NewRoutes(mux, c, g, l)
}
