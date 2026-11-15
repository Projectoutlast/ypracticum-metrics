package metric

import (
	"context"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
	"github.com/Projectoutlast/ypracticum-metrics/internal/repository"
	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
)

type UseCase struct {
	repo repository.Metric
}

func New(r repository.Metric) service.Metric {
	return &UseCase{repo: r}
}

func (uc *UseCase) Get(ctx context.Context) map[string]*models.Metrics {
	return uc.repo.GetAll(ctx)
}
