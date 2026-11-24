package metric

import (
	"context"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (*models.Metrics, error)
	GetAll(ctx context.Context) map[string]*models.Metrics
	Update(ctx context.Context, metric *models.Metrics)
}

type UseCase struct {
	repo Repository
}

func New(r Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) Get(ctx context.Context) map[string]*models.Metrics {
	return uc.repo.GetAll(ctx)
}

func (uc *UseCase) GetByID(ctx context.Context, id string) (*models.Metrics, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *UseCase) Update(ctx context.Context, metric *models.Metrics) {
	uc.repo.Update(ctx, metric)
}
