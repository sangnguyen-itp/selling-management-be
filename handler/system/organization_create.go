package system

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/context"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/error_code"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/service"
)

// OrganizationCreate docs
// @Summary      OrganizationCreate
// @Description  /api/v1/system/organization/create
// @Tags         System
// @Accept       json
// @Produce      json
// @Param   	 body  body   service.OrganizationCreateRequest  true "body"
// @Success      201  {object}  service.OrganizationCreateReply
// @Security     ApiKeyAuth
// @Router       /v1/system/product/create [post]
func OrganizationCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request service.OrganizationCreateRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			app.Response(ctx, 400, error_code.ErrorRequest, nil)
			return
		}

		actor := context.NewBase(ctx)
		request.CreatedAt = actor.UpdateTime
		request.CreatedBy = actor.UserID
		request.UpdatedAt = actor.UpdateTime
		request.UpdatedBy = actor.UserID

		reply, err := service.OrganizationCreate(&request)
		if err != nil {
			logger.Log().Error(domain.OrganizationDomain, "service.OrganizationCreate", err)
			app.Response(ctx, 500, error_code.ServiceError, nil)
			return
		}

		app.Response(ctx, 201, "OK", reply)
	}
}
