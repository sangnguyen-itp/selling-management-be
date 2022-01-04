package common

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/defined/domain"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/service"
)

// Login docs
// @Summary      Login
// @Description  /api/auth/login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param   	 body  body   service.LoginRequest  true "body"
// @Success      200  {object}  service.LoginReply
// @Router       /auth/login [post]
func Login() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var request service.LoginRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			app.Response(ctx, 400, "Bad request", nil)
			return
		}

		reply, err := service.Login(&request)
		if err != nil {
			logger.Log().Error(domain.AuthDomain, "service.Login", err)
			app.Response(ctx, 400, err.Error(), nil)
			return
		}

		app.Response(ctx, 200, "OK", reply)
	}
}
