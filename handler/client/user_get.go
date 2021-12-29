package client

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/context"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/error_code"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/service"
)

// UserGet docs
// @Summary      UserGet
// @Description  /api/v1/user/get
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param   	 body  body   service.UserGetRequest  true "body"
// @Success      201  {object}  service.UserGetReply
// @Security     ApiKeyAuth
// @Router       /v1/client/user/get [post]
func UserGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request service.UserGetRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			app.Response(ctx, 400, error_code.ErrorRequest, nil)
			return
		}

		actor := context.NewBase(ctx)
		request.OrganizationID = actor.OrganizationID

		reply, err := service.UserGet(&request)
		if err != nil {
			logger.Log().Error(domain.UserDomain, "service.GetUser", err)
			app.Response(ctx, 500, error_code.ServiceError, nil)
			return
		}

		app.Response(ctx, 200, "OK", reply)
	}
}
