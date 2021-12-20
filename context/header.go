package context

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetAuthorizeHeader(ctx *gin.Context) string {
	 authorization := ctx.GetHeader("Authorization")
	 authorizations := strings.Split(authorization, " ")
	 if len(authorizations) == 2 {
	 	return authorizations[1]
	 }
	 return ""
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

