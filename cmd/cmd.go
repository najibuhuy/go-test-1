package cmd

import (
	"fmt"
	"go-test-1/delivery/container"
	"go-test-1/delivery/http"
	"go-test-1/infrastructure/logger"
	"go-test-1/infrastructure/shared/constant"
)

func Execute() {
	// start init container
	container := container.SetupContainer()

	// start queue service
	// queue.StartQueueServices(container)

	// start http service
	http := http.ServeHttp(container)
	err := http.Listen(fmt.Sprintf(":%d", container.EnvironmentConfig.App.Port))
	if err != nil {
		// Handle the error, e.g., log it or exit the application
		fmt.Printf("Error starting HTTP server: %s\n", err)
		// Optionally, you can log the error or exit the application.

		logger.LogError(constant.ErrHttpServer, constant.ErrHttpServer, err.Error())

	}
}
