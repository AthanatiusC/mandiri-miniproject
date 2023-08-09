package helper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user *entity.User) (*string, error) {
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

func HandleErrorResponse(gctx *gin.Context, err error) {
	fmt.Printf("[%s%s - %s]Error occured %s\n", gctx.Request.Host, gctx.Request.URL.Path, gctx.ClientIP(), err)
	gctx.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}
