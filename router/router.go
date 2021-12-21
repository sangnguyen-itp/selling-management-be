package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"selling-management-be/conf"
	"selling-management-be/handler"
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
		auth.POST("/login", handler.Login())
		auth.POST("/forgot-password")
	}

	v1 := root.Group("/v1")
	{
		v1.Use(middleware.AuthMiddleware())
		user := v1.Group("/user")
		{
			user.POST("/get", handler.UserGet())
			user.POST("/list", handler.UserList())
		}

		product := v1.Group("/product")
		{
			product.POST("/get", handler.ProductGet())
			product.POST("/list", handler.ProductList())
			product.POST("/create", handler.ProductCreate())
			product.POST("/update", handler.ProductUpdate())
		}
	}

	route.Run(fmt.Sprintf(":%s", conf.EnvConfig.AppPort))
}
