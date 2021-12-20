package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"selling-management-be/conf"
	"time"
)

type jwtHelper struct {
	claims         jwt.MapClaims
	accessSecret   string
	expiryDuration time.Duration
	expiryTime     int64
}

func newJWTHelper() *jwtHelper {
	return &jwtHelper{
		accessSecret:   conf.EnvConfig.SecurityAccessSecret,
		expiryDuration: conf.EnvConfig.SecurityExpiry,
		claims:         make(map[string]interface{}),
		expiryTime:     time.Now().Add(conf.EnvConfig.SecurityExpiry).Unix(),
	}
}

func (j *jwtHelper) Generate(metadata map[string]interface{}) (string, error) {
	j.setDefaultMetadata()
	j.parseMetadataToMapClaims(metadata)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.claims)
	return token.SignedString([]byte(j.accessSecret))
}

func (j *jwtHelper) parseMetadataToMapClaims(metadata Metadata) {
	for key, value := range metadata {
		j.claims[key] = value
	}
}

func (j *jwtHelper) Validate(tokenStr string) (*jwt.Token,bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.accessSecret), nil
	})

	return token, err == nil
}

func (j *jwtHelper) Extract(tokenStr string) (map[string]interface{}, error) {
	token, ok := j.Validate(tokenStr)
	if ok && token.Valid {
		claims, valid := token.Claims.(jwt.MapClaims)
		if valid {
			return claims, nil
		}
	}
	return nil, nil
}

func (j *jwtHelper) setDefaultMetadata() {
	j.claims["exp"] = j.expiryTime
	j.claims["authenticated"] = true
}

func (j *jwtHelper) getExpiryTime() int64 {
	return j.expiryTime
}
