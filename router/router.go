package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"selling-management-be/conf"
	"selling-management-be/handler/client"
	"selling-management-be/handler/common"
	"selling-management-be/handler/system"
	"selling-management-be/middleware"
)

func Run() {
	route := gin.New()
	route.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/user/dns_tracking/"),
		gin.Recovery(),
	)

	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(middleware.CORS())

	route.Static("/public", "./public")

	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	root := route.Group("/api")

	auth := root.Group("/auth")
	{
		auth.POST("/login", common.Login())
		auth.POST("/forgot-password")
	}

	v1 := root.Group("/v1")
	{
		v1.Use(middleware.AuthMiddleware())
		clientEndpoint := v1.Group("/client")
		{
			clientEndpoint.Use(middleware.ClientMiddleware())
			user := clientEndpoint.Group("/user")
			{
				user.POST("/get", client.UserGet())
				user.POST("/list", client.UserList())
			}

			organization := clientEndpoint.Group("/organization")
			{
				organization.POST("/get", client.OrganizationGet())
			}

			product := clientEndpoint.Group("/product")
			{
				product.POST("/get", client.ProductGet())
				product.POST("/list", client.ProductList())
				product.POST("/create", client.ProductCreate())
				product.POST("/update", client.ProductUpdate())
				product.POST("/import", client.ProductImport())
			}
		}

		systemEndpoint := v1.Group("/system")
		{
			clientEndpoint.Use(middleware.SystemMiddleware())
			user := systemEndpoint.Group("/user")
			{
				user.POST("/get", system.UserGet())
				user.POST("/list", client.UserList())
			}

			organization := systemEndpoint.Group("/organization")
			{
				organization.POST("/get", system.OrganizationGet())
				organization.POST("/list", system.OrganizationList())
			}

			product := systemEndpoint.Group("/product")
			{
				product.POST("/get", client.ProductGet())
				product.POST("/list", system.ProductList())
				product.POST("/create", system.ProductCreate())
				product.POST("/update", system.ProductUpdate())
			}
		}
	}

	route.Run(fmt.Sprintf(":%s", conf.EnvConfig.AppPort))
}
