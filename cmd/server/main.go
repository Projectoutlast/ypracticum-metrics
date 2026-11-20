package main

import (
	"flag"

	"github.com/Projectoutlast/ypracticum-metrics/internal/app"
	"github.com/Projectoutlast/ypracticum-metrics/internal/config"
)

func main() {
	addr := flag.String("a", "localhost:8080", "адрес HTTP сервера")

	flag.Parse()

	cfg := config.New()
	cfg.HTTP.Addr = *addr

	app.RunServer(cfg)
}
