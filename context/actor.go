package context

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Actor struct {
	IP             string
	UserID         string
	OrganizationID string
	IsSystem       bool
	UpdateTime     time.Time
}

func NewBase(ctx *gin.Context) *Actor {
	userID, orgID, isSystem := GetActorHeader(ctx)
	return &Actor{
		UserID:         userID,
		IP:             ctx.ClientIP(),
		UpdateTime:     time.Now(),
		OrganizationID: orgID,
		IsSystem:       isSystem,
	}
}
