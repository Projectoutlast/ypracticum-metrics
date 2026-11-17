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
	r.Mu.Lock()
	defer r.Mu.Unlock()

	m, ok := r.DB[id]
	if !ok {
		return nil, models.ErrMetricNotFound
	}

	return m, nil
}

func (r *Repo) GetAll(ctx context.Context) map[string]*models.Metrics {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	res := r.DB

	return res
}

func (r *Repo) Update(ctx context.Context, metric *models.Metrics) {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	m, ok := r.DB[metric.ID]
	if !ok {
		r.DB[metric.ID] = metric

		return
	}

	switch m.MType {
	case models.Counter:
		if m != nil {
			res := *m.Value + *metric.Value

			m.Value = new(res)
		}
	case models.Gauge:
		if m != nil {
			m.Value = metric.Value
		}
	}

	r.DB[metric.ID] = m
}
