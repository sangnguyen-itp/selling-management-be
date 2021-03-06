package service

import (
	"github.com/mitchellh/mapstructure"
	"selling-management-be/model"
	"time"
)

type UserGetRequest struct {
	ID             string `json:"id"`
	OrganizationID string `json:"-"`
}

type UserGetReply struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`

	Username       string    `json:"username"`
	Password       string    `json:"-"`
	Role           string    `json:"role"`
	LoginTime      time.Time `json:"login_time"`
	Status         string    `json:"status"`
	OrganizationID string    `json:"organization_id"`
	IsSystem       bool      `json:"is_system"`
}

func UserGet(request *UserGetRequest) (reply *UserGetReply, err error) {
	var user model.User

	sqlDB := mainService.db
	if len(request.OrganizationID) > 0 {
		sqlDB = sqlDB.Where("organization_id = ?", request.OrganizationID)
	}

	if err = sqlDB.First(&user, "id = ?", request.ID).Error; err != nil {
		return nil, err
	}
	reply, err = toUserGetReply(&user)
	return
}

func toUserGetReply(user *model.User) (*UserGetReply, error) {
	var reply UserGetReply
	if err := mapstructure.Decode(user, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
