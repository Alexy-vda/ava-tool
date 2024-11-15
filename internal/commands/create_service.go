package commands

import (
	"go-init-service/internal/generators"
	"go-init-service/internal/helpers"
)

func CreateServiceCommand() {

	serviceName := helpers.PromptWithDefault("Enter the service name", "new-tool")
	includeDB := helpers.PromptYesNo("Do you want to include a PostgreSQL database?")
	port := helpers.PromptWithDefault("Enter the port", "8080")

	generators.GenerateService(serviceName, includeDB, port)
	helpers.ExecGoModInit(serviceName, serviceName)
}
