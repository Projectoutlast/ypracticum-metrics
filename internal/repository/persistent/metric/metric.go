package metric

import (
	"context"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
	"github.com/Projectoutlast/ypracticum-metrics/internal/repository"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/storage"
)

type Repo struct {
	*storage.InMemory
}

func New(s *storage.InMemory) repository.Metric {
	return &Repo{s}
}

func (r *Repo) GetByID(ctx context.Context, id string) (*models.Metrics, error) {
	return nil, nil
}

func (r *Repo) GetAll(ctx context.Context) map[string]*models.Metrics {
	return r.SelectRows(ctx)
}
