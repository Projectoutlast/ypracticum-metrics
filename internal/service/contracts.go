package service

import "context"

type (
	Counter interface {
		Update(ctx context.Context, value int64) error
	}

	Gauge interface {
		Update(ctx context.Context, value float64) error
	}
)
