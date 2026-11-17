package repository

import (
	"context"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

type (
	Metric interface {
		GetByID(ctx context.Context, id string) (*models.Metrics, error)
		GetAll(ctx context.Context) map[string]*models.Metrics
		Update(ctx context.Context, metric *models.Metrics)
	}

	Collector interface {
		Collect(ctx context.Context, metrics map[string]*models.Metrics)
	}
)
