package middleware

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/context"
	"selling-management-be/defined"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/pkg/token"
	"selling-management-be/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenData := context.GetAuthorizeHeader(ctx)
		if !token.Validate(tokenData) {
			logger.Log().Error(defined.TokenDomain, "token.Validate", nil)
			app.Response(ctx, 401, "invalid token", nil)
			ctx.Abort()
		}

		_, metadataToken, err := token.ExtractMetadata(tokenData)
		if err != nil {
			logger.Log().Error(defined.TokenDomain, "token.ExtractMetadata", err)
			app.Response(ctx, 401, "invalid token", nil)
			ctx.Abort()
		}

		user, err := service.GetUser(&service.UserGetRequest{ID: metadataToken.UserID})
		if err != nil {
			logger.Log().Error(defined.UserDomain, "service.GetUser", err)
			app.Response(ctx, 401, "invalid token", nil)
			ctx.Abort()
		}

		if !acceptUser(user) {
			logger.Log().Error(defined.UserDomain, "acceptUser", nil)
			app.Response(ctx, 401, "invalid token", nil)
			ctx.Abort()
		}

		context.SetActorHeader(ctx, user.ID, user.Role)
		ctx.Next()
	}
}

func acceptUser(user *service.UserGetReply) bool {
	return user.Status == defined.Active
}
