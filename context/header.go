package context

import (
	"errors"
	"github.com/gin-gonic/gin"
	"selling-management-be/defined/error_code"
	"strconv"
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

func GetIsSystem(ctx *gin.Context) string {
	return ctx.Request.Header.Get("is_system")
}

func SetIsSystem(ctx *gin.Context, isSystem bool) {
	ctx.Request.Header.Set("is_system", strconv.FormatBool(isSystem))
}

func SetActorHeader(ctx *gin.Context, userID, organizationID string, isSystem bool) {
	ctx.Request.Header.Set("user_id", userID)
	ctx.Request.Header.Set("organization_id", organizationID)
	ctx.Request.Header.Set("is_system", strconv.FormatBool(isSystem))
}

func GetActorHeader(ctx *gin.Context) (userID, organizationID string, isSystem bool) {
	userID = ctx.Request.Header.Get("user_id")
	organizationID = ctx.Request.Header.Get("organization_id")
	isSystem, _ = strconv.ParseBool(ctx.Request.Header.Get("is_system"))
	return
}

