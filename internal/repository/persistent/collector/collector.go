package collector

import (
	"context"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/storage"
)

type Repo struct {
	*storage.InMemory
}

func New(s *storage.InMemory) *Repo {
	return &Repo{s}
}

func (r *Repo) Collect(ctx context.Context, metrics map[string]*models.Metrics) {
	r.Lock()
	defer r.Unlock()

	for k, v := range metrics {
		if _, ok := r.DB[k]; !ok {
			r.DB[k] = v

			continue
		}

		curr := r.DB[k]

		switch curr.MType {
		case models.Counter:
			if curr.Delta != nil && v.Delta != nil {
				res := *curr.Delta + *v.Delta

				curr.Delta = new(res)
			}
		case models.Gauge:
			curr.Value = v.Value
		}

		r.DB[k] = curr
	}
}

func (r *Repo) Consume(ctx context.Context, id string, sent int64) {
	r.Lock()
	defer r.Unlock()

	curr, ok := r.DB[id]
	if !ok || curr.Delta == nil {
		return
	}

	res := *curr.Delta - sent
	if res < 0 {
		res = 0
	}

	curr.Delta = new(res)
}
