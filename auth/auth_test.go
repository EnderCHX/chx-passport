package auth

import (
	"chx-passport/models/user"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestToken(t *testing.T) {
	user := user.User{
		Username: "test",
	}
	token, err := GetToken(user, "123456", time.Second*10)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("token: %v\n", token)
}

func TestVerify(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJpc3MiOiJjaHhjLmNjIiwiZXhwIjoxNzM3MjY3OTY5LCJpYXQiOjE3MzcyNjc5NTl9.JfGSqZF2QbS-C45XpHOwgal0uCBQNyfUJ8wWc6fMngg"
	secretKey := "123456"
	claims, err := VerifyToken(token, secretKey)
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			fmt.Println("invalid signature")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			fmt.Println("token expired")
		} else {
			fmt.Println("unknown error")
		}
	} else {
		fmt.Println("username:", claims.Username)
		fmt.Println("issuer :", claims.Issuer)
		fmt.Println("issuedat:", claims.IssuedAt.Unix())
		fmt.Println("expires :", claims.ExpiresAt.Unix())
	}
}
