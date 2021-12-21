package service

import (
	"github.com/mitchellh/mapstructure"
	"selling-management-be/model"
	"selling-management-be/pkg/cipher"
	"time"
)

type UserCreateRequest struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	PhoneNumber    string    `json:"phone_number"`
	Address        string    `json:"address"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	Password       string    `json:"-"`
	Role           string    `json:"role"`
	LoginTime      time.Time `json:"login_time"`
	Status         string    `json:"status"`
	OrganizationID string    `json:"organization_id"`

	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

type UserCreateReply struct {
	ID string `json:"id"`
}

func UserCreate(request *UserCreateRequest) (reply *UserCreateReply, err error) {
	var user model.User
	if err = mapstructure.Decode(request, &user); err != nil {
		return nil, err
	}

	user.Password, err = cipher.Hash(user.Password)

	if err = mainService.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &UserCreateReply{ID: user.ID}, nil
}
