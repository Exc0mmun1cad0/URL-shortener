package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	config := config.MustLoad()
	fmt.Println("loaded configuration")
	_ = config

	// TODO: logger initialization

	// TODO: storage initialization

	// TODO: cache initialization

	// TODO: creating router
	// TODO: registering handlers
	// TODO: adding middlewares to router

	// TODO: HTTP server initialization
	// TODO: graceful shutdownm
}
