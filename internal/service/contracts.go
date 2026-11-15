package service

import (
	"context"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

type (
	Counter interface {
		Update(ctx context.Context, value int64) error
	}

	Gauge interface {
		Update(ctx context.Context, value float64) error
	}

	Collector interface {
		Collect(ctx context.Context) error
	}

	Metric interface {
		Get(ctx context.Context) map[string]*models.Metrics
	}
)
