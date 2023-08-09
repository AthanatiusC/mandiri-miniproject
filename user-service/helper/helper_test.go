package helper

import (
	"testing"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/model"
)

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT(&model.User{
		ID:          1,
		Username:    "username",
		AccessLevel: 1,
		Status:      "active",
	})
	if err != nil {
		t.Errorf("error generate jwt: %v", err)
	}
	t.Logf("token: %v", *token)
}
