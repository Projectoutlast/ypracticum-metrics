package agent

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Projectoutlast/ypracticum-metrics/internal/agent"
	"github.com/Projectoutlast/ypracticum-metrics/internal/config"
	collectorRepo "github.com/Projectoutlast/ypracticum-metrics/internal/repository/persistent/collector"
	metricRepo "github.com/Projectoutlast/ypracticum-metrics/internal/repository/persistent/metric"
	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
	collectorSvc "github.com/Projectoutlast/ypracticum-metrics/internal/service/collector"
	metricSvc "github.com/Projectoutlast/ypracticum-metrics/internal/service/metric"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/httpclient"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/storage"
)

type useCases struct {
	c service.Collector
	m service.Metric
}

func RunAgent(cfg *config.Config) {
	l := logger.New("debug")

	s := storage.New()

	uc := initUseCases(s)

	client := httpclient.New()

	a := agent.New(cfg, client, l, uc.c, uc.m)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	a.Run(ctx)
}

func initUseCases(s *storage.InMemory) *useCases {
	collStore := collectorRepo.New(s)
	metricStore := metricRepo.New(s)

	return &useCases{
		c: collectorSvc.New(collStore),
		m: metricSvc.New(metricStore),
	}
}
