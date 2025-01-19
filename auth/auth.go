package auth

import (
	"chx-passport/models/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTPayload struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GetToken(u user.User, secretKey string, expiresAt time.Duration) (string, error) {
	jwtPayload := JWTPayload{
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "chxc.cc",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

func VerifyToken(tokenString string, secretKey string) (*JWTPayload, error) {
	t, err := jwt.ParseWithClaims(tokenString, &JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := t.Claims.(*JWTPayload); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
