package v1

import (
	"context"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
)

type Metric interface {
	Get(ctx context.Context) map[string]*models.Metrics
	GetByID(ctx context.Context, id string) (*models.Metrics, error)
	Update(ctx context.Context, metric *models.Metrics)
}

type V1 struct {
	m Metric
	l logger.Interface
}
