package helper

import (
	"time"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/model"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *model.User) (*string, error) {
	exp := time.Now().Add(time.Minute * 1).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": exp,
	})

	token, err := claims.SignedString([]byte("test"))
	if err != nil {
		return nil, err
	}
	return &token, nil
}
