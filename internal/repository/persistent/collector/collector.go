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
	r.Mu.Lock()

	for k, v := range metrics {
		if _, ok := r.DB[k]; !ok {
			r.DB[k] = v

			continue
		}

		curr := r.DB[k]

		switch curr.MType {
		case models.Counter:
			if curr != nil {
				res := *curr.Value + *v.Value

				curr.Value = new(res)
			}
		case models.Gauge:
			if curr != nil {
				curr.Value = v.Value
			}
		}

		r.DB[k] = curr
	}

	r.Mu.Unlock()
}
