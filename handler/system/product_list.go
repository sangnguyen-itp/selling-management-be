package system

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/error_code"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/service"
)

// ProductList docs
// @Summary      ProductList
// @Description  /api/v1/product/list
// @Tags         System
// @Accept       json
// @Produce      json
// @Param   	 body  body   service.ProductListRequest  true "body"
// @Success      201  {object}  []service.ProductGetReply
// @Security     ApiKeyAuth
// @Router       /v1/system/product/list [post]
func ProductList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request service.ProductListRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			app.Response(ctx, 400, error_code.ErrorRequest, nil)
			return
		}

		reply, err := service.ProductList(&request)
		if err != nil {
			logger.Log().Error(domain.ProductDomain, "service.UserList", err)
			app.Response(ctx, 400, error_code.ServiceError, nil)
			return
		}

		app.Response(ctx, 200, "OK", reply)
	}
}

