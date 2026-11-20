package service

import (
	"context"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

type (
	Collector interface {
		Collect(ctx context.Context) error
	}

	Metric interface {
		Get(ctx context.Context) map[string]*models.Metrics
		GetByID(ctx context.Context, id string) (*models.Metrics, error)
		Update(ctx context.Context, metric *models.Metrics)
	}
)
