package main

import (
	"github.com/amorimdotnetdevops-debug/go-user-control/config"
	"github.com/amorimdotnetdevops-debug/go-user-control/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Failed to initialize config: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
