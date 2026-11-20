package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
	"github.com/go-chi/chi/v5"
)

const (
	_defaultAddr            = ":80"
	_defaultReadTimeout     = 15 * time.Second
	_defaultWriteTimeout    = 15 * time.Second
	_defaultShutdownTimeout = 10 * time.Second
)

type Server struct {
	Mux *chi.Mux

	ctx context.Context
	eg  *errgroup.Group

	httpServer *http.Server
	notify     chan error

	addr            string
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration

	logger logger.Interface
}

func New(l logger.Interface, opts ...Option) *Server {
	group, ctx := errgroup.WithContext(context.Background())
	group.SetLimit(1)

	s := &Server{
		ctx:             ctx,
		eg:              group,
		notify:          make(chan error, 1),
		addr:            _defaultAddr,
		readTimeout:     _defaultReadTimeout,
		writeTimeout:    _defaultWriteTimeout,
		shutdownTimeout: _defaultShutdownTimeout,
		logger:          l,
	}

	for _, opt := range opts {
		opt(s)
	}

	m := chi.NewRouter()

	s.Mux = m
	s.httpServer = &http.Server{
		Addr:         s.addr,
		Handler:      m,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
	}

	return s
}

func (s *Server) Start() {
	s.eg.Go(func() error {
		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.notify <- err

			close(s.notify)

			return err
		}

		return nil
	})

	s.logger.Info(fmt.Sprintf("restapi server - Server - Started on %s", s.addr))
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	var shutdownErrs []error

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	err := s.httpServer.Shutdown(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		s.logger.Error("restapi server - Server - Shutdown - s.httpServer.Shutdown", "error", err)
		shutdownErrs = append(shutdownErrs, err)
	}

	err = s.eg.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		s.logger.Error("restapi server - Server - Shutdown - s.eg.Wait", "error", err)
		shutdownErrs = append(shutdownErrs, err)
	}

	s.logger.Info("restapi server - Server - Shutdown")

	return errors.Join(shutdownErrs...)
}
