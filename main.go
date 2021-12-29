package main

import (
	"selling-management-be/defined/domain"
	"selling-management-be/docs"
	"selling-management-be/init"
	"selling-management-be/pkg/logger"
	"selling-management-be/router"
	"selling-management-be/service"
)

// @title           Selling Management
// @version         1.0
// @description     API Management
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      selling.management
// @BasePath  /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logger.Log().Info(domain.SystemDomain, "Starting service ...")
	err := service.NewService()
	if err != nil {
		logger.Log().Error(domain.SystemDomain, "service.NewService", err)
	}

	docs.SetupSwaggerInfo()
	init.InitDefaultData()

	router.Run()
}
