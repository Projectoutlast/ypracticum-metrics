package counter

import (
	"context"

	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
)

type UseCase struct{}

func New() service.Counter {
	return &UseCase{}
}

func (uc *UseCase) Update(ctx context.Context, value int64) error {
	return nil
}
