package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type JwtCustomClaims struct {
	ID uint
	jwt.RegisteredClaims
}

func GeneratorToken(id uint) (string, error) {
	jwtCustomClaims := JwtCustomClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		}}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCustomClaims)
	return jwtToken.SignedString([]byte(viper.GetString("jwt.signingKey")))
}

func ParseToken(tokenString string) (JwtCustomClaims, error) {
	jwtCustomClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt.signingKey")), nil
	})

	if err == nil && !token.Valid {
		err = errors.New("invalid token")
	}
	return jwtCustomClaims, err
}

// token 是否是正确的，true 正确，false 错误
func TokenValid(tokenString string) bool {
	_, err := ParseToken(tokenString)
	if err != nil {
		return false
	}
	return true
}
