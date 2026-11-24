package httpclient

import (
	"net/http"
	"time"
)

const _defaultTimeout = 5 * time.Second

type HTTPClient struct {
	Client *http.Client

	timeout time.Duration
}

func New(opts ...Option) *HTTPClient {
	c := &HTTPClient{
		timeout: _defaultTimeout,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Client = &http.Client{
		Timeout: c.timeout,
	}

	return c
}
