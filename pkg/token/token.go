package token

import (
	"errors"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"selling-management-be/defined"
	"selling-management-be/pkg/generate_id"
)

type AuthData struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	IP     string `json:"ip"`
}

type Metadata map[string]interface{}

type Token struct {
	TokenKey   string
	TokenValue string
	ExpiryTime int64
}

func Generate(authData *AuthData) (*Token, error) {
	if authData == nil {
		return nil, errors.New("invalid metadata")
	}

	jwt := newJWTHelper()
	tokenKey := generate_id.NewID(defined.SystemTokenDomain)
	tokenValue, err := jwt.Generate(structs.Map(authData))
	if err != nil {
		return nil,  err
	}

	return &Token{
		TokenKey:   tokenKey,
		TokenValue: tokenValue,
		ExpiryTime: jwt.getExpiryTime(),
	}, nil
}

func Validate(token string) bool {
	jwt := newJWTHelper()
	_, valid := jwt.Validate(token)
	return valid
}

func ExtractMetadata(token string) (*AuthData, error) {
	jwt := newJWTHelper()
	metadata, err := jwt.Extract(token)
	if err != nil {
		return nil, err
	}

	var authData AuthData
	err = mapstructure.Decode(metadata, &authData)
	if err != nil {
		return nil, err
	}

	return &authData, nil
}