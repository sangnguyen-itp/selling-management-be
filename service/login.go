package service

import (
	"errors"
	"selling-management-be/helper/pagination"
	"selling-management-be/pkg/cipher"
	"selling-management-be/pkg/token"
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	AccessToken string `json:"access_token"`
	OrganizationID string `json:"organization_id"`
	UserID string `json:"user_id"`
	LoginTime time.Time `json:"login_time"`
	IsSystem bool `json:"is_system"`
}

func Login(request *LoginRequest) (reply *LoginReply, err error) {
	userList, err := UserList(&UserListRequest{
		Usernames: []string{request.Username},
		Pagination: pagination.Pagination{
			Limit: 1,
		},
	})
	if err != nil {
		return nil, err
	}

	if len(userList) < 1 {
		return nil, errors.New("username is not found")
	}

	userData := userList[0]
	if !cipher.VerifyPassword(userData.Password, request.Password) {
		return nil, errors.New("password is incorrect")
	}

	tokenData, err := token.Generate(&token.AuthData{
		UserID: userData.ID,
	})

	return &LoginReply{
		AccessToken: tokenData.TokenValue,
		OrganizationID: userData.OrganizationID,
		UserID: userData.ID,
		LoginTime: time.Now(),
		IsSystem: userData.IsSystem,
	}, nil
}
