package middleware

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/context"
	"selling-management-be/defined/domain"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"strconv"
)

func ClientMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isSystemHeader := context.GetIsSystem(ctx)
		isSystem, err := strconv.ParseBool(isSystemHeader)
		if err != nil {
			logger.Log().Error(domain.AuthDomain, "GetIsSystem", err)
			app.Response(ctx, 403, "access denied", nil)
			ctx.Abort()
			return
		}
		if isSystem {
			logger.Log().Error(domain.AuthDomain, "GetIsSystem", err)
			app.Response(ctx, 403, "access denied", nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
