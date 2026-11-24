package main

import (
	"github.com/Projectoutlast/ypracticum-metrics/internal/app"
	"github.com/Projectoutlast/ypracticum-metrics/internal/config"
)

func main() {
	app.RunServer(config.GetServerConfig())
}
