package metric

import (
	"context"
	"maps"
	"sync"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

type Repo struct {
	db map[string]*models.Metrics
	mu sync.Mutex
}

func New() *Repo {
	return &Repo{db: make(map[string]*models.Metrics)}
}

func (r *Repo) GetByID(ctx context.Context, id string) (*models.Metrics, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	m, ok := r.db[id]
	if !ok {
		return nil, models.ErrMetricNotFound
	}

	return m, nil
}

func (r *Repo) GetAll(ctx context.Context) map[string]*models.Metrics {
	r.mu.Lock()
	defer r.mu.Unlock()

	res := make(map[string]*models.Metrics, len(r.db))

	maps.Copy(res, r.db)

	return res
}

func (r *Repo) Update(ctx context.Context, metric *models.Metrics) {
	r.mu.Lock()
	defer r.mu.Unlock()

	m, ok := r.db[metric.ID]
	if !ok {
		r.db[metric.ID] = metric

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

	r.db[metric.ID] = m
}

func (r *Repo) Collect(ctx context.Context, metrics map[string]*models.Metrics) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for k, v := range metrics {
		if _, ok := r.db[k]; !ok {
			r.db[k] = v

			continue
		}

		curr := r.db[k]

		switch curr.MType {
		case models.Counter:
			if curr.Delta != nil && v.Delta != nil {
				res := *curr.Delta + *v.Delta

				curr.Delta = new(res)
			}
		case models.Gauge:
			curr.Value = v.Value
		}

		r.db[k] = curr
	}
}

func (r *Repo) Consume(ctx context.Context, id string, sent int64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	curr, ok := r.db[id]
	if !ok || curr.Delta == nil {
		return
	}

	res := max(*curr.Delta-sent, 0)

	curr.Delta = new(res)
}
