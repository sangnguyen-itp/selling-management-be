package context

import (
	"github.com/gin-gonic/gin"
)

func GetAuthorizeHeader(ctx *gin.Context) string {
	 return ctx.GetHeader("Authorization")
}

func GetClientIP(ctx *gin.Context) string {
	return ctx.ClientIP()
}
