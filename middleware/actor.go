package middleware

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/context"
	"selling-management-be/defined"
	"selling-management-be/model"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/pkg/token"
	"selling-management-be/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenData := context.GetAuthorizeHeader(ctx)
		if token.Validate(tokenData) {
			logger.Log().Error(defined.SystemTokenDomain, "token.Validate", nil)
			app.Response(ctx, 401, "invalid token", nil)
			return
		}

		metadataToken, err := token.ExtractMetadata(tokenData)
		if err != nil {
			logger.Log().Error(defined.SystemTokenDomain, "token.ExtractMetadata", err)
			app.Response(ctx, 401, "invalid token", nil)
			return
		}

		user, err := service.GetUser(&service.UserGetRequest{ID: metadataToken.UserID})
		if err != nil {
			logger.Log().Error(defined.UserDomain, "service.GetUser", err)
			app.Response(ctx, 401, "invalid token", nil)
			return
		}

		if !acceptUser(user) {
			logger.Log().Error(defined.UserDomain, "acceptUser", nil)
			app.Response(ctx, 401, "invalid token", nil)
			ctx.Abort()
		}

		ctx.Next()
	}
}

func acceptUser(user *model.User) bool {
	return user.Status == defined.Active
}
