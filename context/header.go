package context

import (
	"errors"
	"github.com/gin-gonic/gin"
	"selling-management-be/defined/error_code"
	"strings"
)

func GetAuthorizeHeader(ctx *gin.Context) (string, error) {
	 authorization := ctx.GetHeader("Authorization")
	 authorizations := strings.Split(authorization, " ")
	 if len(authorizations) == 2 {
	 	return authorizations[1], nil
	 }

	 return "", errors.New(error_code.AuthorizationHeaderNull)
}

func GetClientIP(ctx *gin.Context) string {
	return ctx.ClientIP()
}

func SetActorHeader(ctx *gin.Context, userID, role string) {
	ctx.Request.Header.Set("user_id", userID)
}

func GetActorHeader(ctx *gin.Context) (userID string) {
	userID = ctx.Request.Header.Get("user_id")
	return
}

