package auth

import (
	"chx-passport/models/user"
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	user := user.User{
		ID:       1,
		Username: "test",
	}
	token, err := GetToken(user, "123456")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("token: %v\n", token)
}

func TestVerify(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3QiLCJpc3MiOiJjaHhjLmNjIiwiZXhwIjoxNzM2ODY1MTk1LCJuYmYiOjE3MzQyNzMxOTUsImlhdCI6MTczNDI3MzE5NX0.81nPnA85VCayYTRdvaqas_l3my-iOr0Jtv9AWGZFsslQ"
	secretKey := "123456"
	claims, err := VerifyToken(token, secretKey)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("claims: %v\n", claims)
}
