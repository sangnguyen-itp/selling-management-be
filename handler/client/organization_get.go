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

// OrganizationGet docs
// @Summary      OrganizationGet
// @Description  /api/v1/system/organization/get
// @Tags         Client
// @Accept       json
// @Produce      json
// @Param   	 body  body   service.OrganizationGetRequest  true "body"
// @Success      201  {object}  []service.OrganizationGetReply
// @Security     ApiKeyAuth
// @Router       /v1/client/organization/get [post]
func OrganizationGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request service.OrganizationGetRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			app.Response(ctx, 400, error_code.ErrorRequest, nil)
			return
		}

		actor := context.NewBase(ctx)
		request.ID = actor.OrganizationID

		reply, err := service.OrganizationGet(&request)
		if err != nil {
			logger.Log().Error(domain.OrganizationDomain, "service.OrganizationGet", err)
			app.Response(ctx, 400, error_code.ServiceError, nil)
			return
		}

		app.Response(ctx, 200, "OK", reply)
	}
}



