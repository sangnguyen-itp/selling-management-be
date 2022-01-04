package system

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/error_code"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/service"
)

// UserList docs
// @Summary      UserList
// @Description  /api/v1/user/get
// @Tags         System
// @Accept       json
// @Produce      json
// @Param   	 body  body   service.UserListRequest  true "body"
// @Success      201  {object}  []service.UserGetReply
// @Security     ApiKeyAuth
// @Router       /v1/system/user/list [post]
func UserList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request service.UserListRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			app.Response(ctx, 400, error_code.ErrorRequest, nil)
			return
		}

		reply, err := service.UserList(&request)
		if err != nil {
			logger.Log().Error(domain.UserDomain, "service.UserList", err)
			app.Response(ctx, 400, error_code.ServiceError, nil)
			return
		}

		app.Response(ctx, 200, "OK", reply)
	}
}
