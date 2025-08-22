package main

import (
	"github.com/amorimdotnetdevops-debug/go-user-control/config"
	"github.com/amorimdotnetdevops-debug/go-user-control/router"
)

func main() {
	// Initialize configs
	err := config.Init()
	if err != nil {
		panic("Failed to initialize config: " + err.Error())
	}

	// Initialize router
	router.Initialize()
}
