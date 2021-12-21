package handler

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/context"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/error_code"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/service"
)

// ProductCreate docs
// @Summary      ProductCreate
// @Description  /api/v1/product/create
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param   	 body  body   service.ProductCreateRequest  true "body"
// @Success      201  {object}  service.ProductCreateReply
// @Security     ApiKeyAuth
// @Router       /v1/product/create [post]
func ProductCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request service.ProductCreateRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			app.Response(ctx, 400, error_code.ErrorRequest, nil)
			return
		}

		actor := context.NewBase(ctx)
		request.CreatedAt = actor.UpdateTime
		request.CreatedBy = actor.UserID
		request.UpdatedAt = actor.UpdateTime
		request.UpdatedBy = actor.UserID
		if !actor.IsSystem {
			request.OrganizationID = actor.OrganizationID
		} else {
			if !ValidateOrganizationID(request.OrganizationID) {
				logger.Log().Error(domain.OrganizationDomain, "ValidateOrganizationID", nil)
				app.Response(ctx, 400, error_code.ErrorRequest, nil)
				return
			}
		}

		reply, err := service.ProductCreate(&request)
		if err != nil {
			logger.Log().Error(domain.ProductDomain, "service.ProductGet", err)
			app.Response(ctx, 500, error_code.ServiceError, nil)
			return
		}

		app.Response(ctx, 201, "OK", reply)
	}
}
