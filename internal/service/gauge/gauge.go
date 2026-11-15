package gauge

import (
	"context"

	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
)

type UseCase struct{}

func New() service.Gauge {
	return &UseCase{}
}

func (uc *UseCase) Update(ctx context.Context, value float64) error {
	return nil
}
