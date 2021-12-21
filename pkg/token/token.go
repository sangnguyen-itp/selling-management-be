package token

import (
	"errors"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"selling-management-be/defined/domain"
	"selling-management-be/pkg/generate_id"
)

type AuthData struct {
	UserID         string `json:"user_id"`
	OrganizationID string `json:"organization_id"`
	IsSystem       bool   `json:"is_system"`
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
	tokenKey := generate_id.NewID(domain.TokenDomain)
	metadata := structs.Map(authData)
	metadata["token_key"] = tokenKey

	tokenValue, err := jwt.Generate(metadata)
	if err != nil {
		return nil, err
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

func ExtractMetadata(token string) (tokenKey string, authData *AuthData, err error) {
	jwt := newJWTHelper()
	metadata, err := jwt.Extract(token)
	if err != nil {
		return "", nil, err
	}

	err = mapstructure.Decode(metadata, &authData)
	if err != nil {
		return "", nil, err
	}

	tokenKey = metadata["token_key"].(string)
	return
}
