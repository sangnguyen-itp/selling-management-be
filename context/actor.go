package context

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Actor struct {
	IP         string
	UserID     string
	UpdateTime time.Time
}

func NewBase(ctx *gin.Context) *Actor {
	userID := GetActorHeader(ctx)
	return &Actor{
		UserID: userID,
		IP: ctx.ClientIP(),
		UpdateTime: time.Now(),
	}
}
