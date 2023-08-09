package controller

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func extractBearerAuth(auth string) (string, error) {
	if auth == "" {
		return "", errors.New("unauthorized")
	}

	jwtToken := strings.Split(auth, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("invalid token")
	}

	return jwtToken[1], nil
}

func (c *Controller) JWTMiddleware(gctx *gin.Context) {
	auth, err := extractBearerAuth(gctx.GetHeader("Authorization"))
	if err != nil {
		gctx.AbortWithStatusJSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(c.Config.SecretConfig.JWT), nil
	})
	if err != nil {
		gctx.AbortWithStatusJSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	if token.Valid {
		gctx.Next()
	}
}
