package config

import (
	"flag"
	"time"
)

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

func GetServerConfig() *Config {
	addr := flag.String("a", "localhost:8080", "адрес HTTP сервера")

	flag.Parse()

	return &Config{
		HTTP: http{
			Addr: *addr,
		},
	}
}

func GetAgentConfig() *Config {
	addr := flag.String("a", "localhost:8080", "адрес HTTP эндпоинта сервера сбора метрик")
	reportInterval := flag.Int("r", 10, "частота отправки метрик на сервер")
	pollInterval := flag.Int("p", 2, "частота опроса метрик")

	flag.Parse()

	return &Config{
		Agent: agent{
			ServerAddr:     *addr,
			ReportInterval: time.Duration(*reportInterval) * time.Second,
			PollInterval:   time.Duration(*pollInterval) * time.Second,
		},
	}
}
