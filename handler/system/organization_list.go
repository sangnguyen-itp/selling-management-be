package system

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/error_code"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/service"
)

// OrganizationList docs
// @Summary      OrganizationList
// @Description  /api/v1/system/organization/list
// @Tags         System
// @Accept       json
// @Produce      json
// @Param   	 body  body   service.OrganizationListRequest  true "body"
// @Success      201  {object}  []service.OrganizationGetReply
// @Security     ApiKeyAuth
// @Router       /v1/system/organization/list [post]
func OrganizationList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request service.OrganizationListRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			app.Response(ctx, 400, error_code.ErrorRequest, nil)
			return
		}

		reply, err := service.OrganizationList(&request)
		if err != nil {
			logger.Log().Error(domain.OrganizationDomain, "service.OrganizationList", err)
			app.Response(ctx, 400, error_code.ServiceError, nil)
			return
		}

		app.Response(ctx, 200, "OK", reply)
	}
}
