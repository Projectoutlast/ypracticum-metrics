package httpserver

import (
	"time"
)

type Option func(*Server)

func WithAddr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

func WithReadTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.readTimeout = t
	}
}

func WithWriteTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.writeTimeout = t
	}
}

func WithShutdownTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = t
	}
}
