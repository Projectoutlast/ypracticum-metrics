package v1

import (
	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
)

type V1 struct {
	m service.Metric
	l logger.Interface
}
