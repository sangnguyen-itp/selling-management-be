package model

import (
	"encoding/json"
	"fmt"
	"selling-management-be/defined"
	"selling-management-be/pkg/logger"
	"time"
)

type User struct {
	ID          string `gorm:"primaryKey"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`

	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	LoginTime time.Time `json:"login_time"`
	Status    string    `json:"status"`

	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(ID string, firstName string, lastName string, phoneNumbers string, addresses string, username string, password string, loginTime time.Time, updatedBy string, updatedAt time.Time, createdBy string, createdAt time.Time) *User {
	return &User{ID: ID, FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumbers, Address: addresses, Username: username, Password: password, LoginTime: loginTime, UpdatedBy: updatedBy, UpdatedAt: updatedAt, CreatedBy: createdBy, CreatedAt: createdAt}
}

func (u *User) FullName() string {
	if u != nil {
		return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	}
	logger.Log().Error(defined.UserDomain, "user is nil", nil)
	return ""
}

func (u *User) Encode() string {
	bytes, err := json.Marshal(u)
	if err != nil {
		logger.Log().Error(defined.UserDomain, "Encode failed", err)
		return ""
	}
	return string(bytes)
}

func (u *User) Decode(bytes string) error {
	err := json.Unmarshal([]byte(bytes), u)
	logger.Log().Error(defined.UserDomain, "Encode failed", err)
	return err
}
