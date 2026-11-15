package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Projectoutlast/ypracticum-metrics/internal/handler/restapi"
	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
	counterSvc "github.com/Projectoutlast/ypracticum-metrics/internal/service/counter"
	gaugeSvc "github.com/Projectoutlast/ypracticum-metrics/internal/service/gauge"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/httpserver"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
)

type useCases struct {
	c service.Counter
	g service.Gauge
}

type server struct {
	http *httpserver.Server
}

func RunServer() {
	l := logger.New("debug")

	uc := initUseCases()

	s := initServer(l, uc)
	s.startServer()

	waitForShutdown(s, l)
}

func initUseCases() *useCases {
	return &useCases{
		c: counterSvc.New(),
		g: gaugeSvc.New(),
	}
}

func initServer(l logger.Interface, uc *useCases) server {
	httpServer := httpserver.New(l,
		httpserver.WithAddr("http://localhost"),
		httpserver.WithPort("8080"),
	)

	restapi.NewRouter(httpServer.Mux, uc.c, uc.g, l)

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
