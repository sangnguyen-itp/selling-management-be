package handler

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/defined"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/service"
)

func UserGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request service.UserGetRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			app.Response(ctx, 400, "Bad request", nil)
			return
		}

		reply, err := service.GetUser(&request)
		if err != nil {
			logger.Log().Error(defined.UserDomain, "service.GetUser", err)
			app.Response(ctx, 500, "Something went wrong", nil)
			return
		}

		app.Response(ctx, 200, "OK", reply)
	}
}
