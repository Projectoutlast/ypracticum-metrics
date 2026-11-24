package main

import (
	"github.com/Projectoutlast/ypracticum-metrics/internal/app/agent"
	"github.com/Projectoutlast/ypracticum-metrics/internal/config"
)

func main() {
	agent.RunAgent(config.GetAgentConfig())
}
