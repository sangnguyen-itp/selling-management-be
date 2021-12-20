package service

import (
	"errors"
	"selling-management-be/helper"
	"selling-management-be/pkg/cipher"
	"selling-management-be/pkg/token"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	AccessToken string
}

func Login(request *LoginRequest) (reply *LoginReply, err error) {
	userList, err := UserList(&UserListRequest{
		Usernames: []string{request.Username},
		Pagination: helper.Pagination{
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

	return &LoginReply{AccessToken: tokenData.TokenValue}, nil
}
