package main

import (
	"flag"
	"time"

	"github.com/Projectoutlast/ypracticum-metrics/internal/app/agent"
	"github.com/Projectoutlast/ypracticum-metrics/internal/config"
)

func main() {
	addr := flag.String("a", "localhost:8080", "адрес HTTP эндпоинта сервера сбора метрик")
	reportInterval := flag.Int("r", 10, "частота отправки метрик на сервер")
	pollInterval := flag.Int("p", 2, "частота опроса метрик")

	flag.Parse()

	cfg := config.New()

	cfg.Agent.ServerAddr = *addr
	cfg.Agent.ReportInterval = time.Duration(*reportInterval) * time.Second
	cfg.Agent.PollInterval = time.Duration(*pollInterval) * time.Second

	agent.RunAgent(cfg)
}
