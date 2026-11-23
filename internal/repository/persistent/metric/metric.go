package metric

import (
	"context"
	"maps"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/storage"
)

type Repo struct {
	*storage.InMemory
}

func New(s *storage.InMemory) *Repo {
	return &Repo{s}
}

func (r *Repo) GetByID(ctx context.Context, id string) (*models.Metrics, error) {
	r.Lock()
	defer r.Unlock()

	m, ok := r.DB[id]
	if !ok {
		return nil, models.ErrMetricNotFound
	}

	return m, nil
}

func (r *Repo) GetAll(ctx context.Context) map[string]*models.Metrics {
	r.Lock()
	defer r.Unlock()

	res := make(map[string]*models.Metrics, len(r.DB))

	maps.Copy(res, r.DB)

	return res
}

func (r *Repo) Update(ctx context.Context, metric *models.Metrics) {
	r.Lock()
	defer r.Unlock()

	m, ok := r.DB[metric.ID]
	if !ok {
		r.DB[metric.ID] = metric

		return
	}

	switch m.MType {
	case models.Counter:
		if m.Delta != nil && metric.Delta != nil {
			res := *m.Delta + *metric.Delta

			m.Delta = new(res)
		}
	case models.Gauge:
		m.Value = metric.Value
	}

	r.DB[metric.ID] = m
}
