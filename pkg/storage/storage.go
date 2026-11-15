package storage

import (
	"context"
	"sync"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

type InMemory struct {
	db map[string]*models.Metrics
	m  sync.Mutex
}

func New() *InMemory {
	return &InMemory{db: make(map[string]*models.Metrics)}
}

func (i *InMemory) SelectRows(ctx context.Context) map[string]*models.Metrics {
	return i.db
}

func (i *InMemory) UpdateAll(ctx context.Context, metrics map[string]*models.Metrics) {
	i.m.Lock()

	for k, v := range metrics {
		if _, ok := i.db[k]; !ok {
			i.db[k] = v

			continue
		}

		curr := i.db[k]

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

		i.db[k] = curr
	}

	i.m.Unlock()
}
