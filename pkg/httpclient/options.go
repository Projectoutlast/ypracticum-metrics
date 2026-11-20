package httpclient

import "time"

type Option func(*HTTPClient)

func WithTimeout(t time.Duration) Option {
	return func(c *HTTPClient) {
		c.timeout = t
	}
}
