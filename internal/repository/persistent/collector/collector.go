package collector

import (
	"context"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
	"github.com/Projectoutlast/ypracticum-metrics/internal/repository"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/storage"
)

type Repo struct {
	*storage.InMemory
}

func New(s *storage.InMemory) repository.Collector {
	return &Repo{s}
}

func (r *Repo) Collect(ctx context.Context, metrics map[string]*models.Metrics) {
	r.UpdateAll(ctx, metrics)
}
