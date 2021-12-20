package main

import (
	"selling-management-be/defined"
	"selling-management-be/pkg/logger"
	"selling-management-be/router"
	"selling-management-be/service"
)

func main() {
	logger.Log().Info(defined.SystemDomain, "Starting service ...")
	err := service.NewService()
	if err != nil {
		logger.Log().Error(defined.SystemDomain, "service.NewService", err)
	}


	router.Run()
}
