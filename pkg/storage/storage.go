package storage

import (
	"sync"

	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
)

type InMemory struct {
	DB map[string]*models.Metrics
	Mu sync.Mutex
}

func New() *InMemory {
	return &InMemory{DB: make(map[string]*models.Metrics)}
}
