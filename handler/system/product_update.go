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

// ProductUpdate docs
// @Summary      ProductUpdate
// @Description  /api/v1/product/update
// @Tags         System
// @Accept       json
// @Produce      json
// @Param   	 body  body   service.ProductUpdateRequest  true "body"
// @Success      201  {object}  service.ProductUpdateReply
// @Security     ApiKeyAuth
// @Router       /v1/system/product/update [post]
func ProductUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request service.ProductUpdateRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			app.Response(ctx, 400, error_code.ErrorRequest, nil)
			return
		}

		actor := context.NewBase(ctx)
		request.UpdatedAt = actor.UpdateTime
		request.UpdatedBy = actor.UserID

		reply, err := service.ProductUpdate(&request)
		if err != nil {
			logger.Log().Error(domain.ProductDomain, "service.ProductUpdate", err)
			app.Response(ctx, 400, error_code.ServiceError, nil)
			return
		}

		app.Response(ctx, 200, "OK", reply)
	}
}

