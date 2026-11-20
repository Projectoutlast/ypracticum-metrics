package config

import "time"

type (
	Config struct {
		HTTP  http
		Agent agent
	}

	http struct {
		Addr string
	}

	agent struct {
		ServerAddr     string
		ReportInterval time.Duration
		PollInterval   time.Duration
	}
)

func New() *Config {
	return &Config{}
}
