package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Projectoutlast/ypracticum-metrics/internal/config"
	"github.com/Projectoutlast/ypracticum-metrics/internal/handler/restapi"
	"github.com/Projectoutlast/ypracticum-metrics/internal/repository/persistent/metric"
	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
	svcMetric "github.com/Projectoutlast/ypracticum-metrics/internal/service/metric"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/httpserver"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/storage"
)

type useCases struct {
	m service.Metric
}

type server struct {
	http *httpserver.Server
}

func RunServer(cfg *config.Config) {
	l := logger.New("debug")

	inMemory := storage.New()

	uc := initUseCases(inMemory)

	s := initServer(cfg, l, uc)
	s.startServer()

	waitForShutdown(s, l)
}

func initUseCases(s *storage.InMemory) *useCases {
	metricRepo := metric.New(s)

	return &useCases{
		m: svcMetric.New(metricRepo),
	}
}

func initServer(cfg *config.Config, l logger.Interface, uc *useCases) server {
	httpServer := httpserver.New(l,
		httpserver.WithAddr(cfg.HTTP.Addr),
	)

	restapi.NewRouter(httpServer.Mux, uc.m, l)

	return server{http: httpServer}
}

func waitForShutdown(s server, l logger.Interface) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-interrupt:
		l.Info("app - Run - signal: ", "signal", sig.String())
	case err := <-s.http.Notify():
		l.Error("app - Run - httpServer.Notify", "error", err)
	}

	s.shutdownServer(l)
}

func (s *server) startServer() {
	s.http.Start()
}

func (s *server) shutdownServer(l logger.Interface) {
	if err := s.http.Shutdown(); err != nil {
		l.Error("app - Run - httpServer.Shutdown", "error", err)
	}
}
