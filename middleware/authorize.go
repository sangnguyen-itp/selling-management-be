package middleware

import (
	"github.com/gin-gonic/gin"
	"selling-management-be/context"
	"selling-management-be/defined/domain"
	"selling-management-be/defined/user_status"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/pkg/token"
	"selling-management-be/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenData, err := context.GetAuthorizeHeader(ctx)
		if err != nil {
			logger.Log().Error(domain.TokenDomain, "context.GetAuthorizeHeader", nil)
			app.Response(ctx, 401, "token not found", nil)
			ctx.Abort()
			return
		}

		if !token.Validate(tokenData) {
			logger.Log().Error(domain.TokenDomain, "token.Validate", nil)
			app.Response(ctx, 401, "invalid token", nil)
			ctx.Abort()
			return
		}

		_, metadataToken, err := token.ExtractMetadata(tokenData)
		if err != nil {
			logger.Log().Error(domain.TokenDomain, "token.ExtractMetadata", err)
			app.Response(ctx, 401, "invalid token", nil)
			ctx.Abort()
			return
		}

		user, err := service.UserGet(&service.UserGetRequest{ID: metadataToken.UserID})
		if err != nil {
			logger.Log().Error(domain.UserDomain, "service.GetUser", err)
			app.Response(ctx, 401, "invalid token", nil)
			ctx.Abort()
			return
		}

		if !acceptUser(user) {
			logger.Log().Error(domain.UserDomain, "acceptUser", nil)
			app.Response(ctx, 401, "invalid token", nil)
			ctx.Abort()
			return
		}

		context.SetActorHeader(ctx, user.ID, user.OrganizationID, user.IsSystem)
		ctx.Next()
	}
}

func acceptUser(user *service.UserGetReply) bool {
	return user.Status == user_status.Active
}
