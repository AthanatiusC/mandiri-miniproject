package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/internal/service"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/model"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c *Controller) GetUsers(gin *gin.Context) {
	var request model.UserRequest
	gin.ShouldBindJSON(&request)
	if request.ID != 0 {
		gin.JSON(http.StatusOK, request)
		return
	}

	gin.JSON(http.StatusOK, request)
}

func (c *Controller) UpdateUser(gin *gin.Context) {
	var request model.UserRequest
	decoder := json.NewDecoder(gin.Request.Body)
	decoder.Decode(&request)

	gin.JSON(http.StatusOK, request)
}

func (c *Controller) CreateUser(gin *gin.Context) {
	var request model.UserRequest
	decoder := json.NewDecoder(gin.Request.Body)
	decoder.Decode(&request)

	gin.JSON(http.StatusOK, request)
}

func (c *Controller) DeleteUser(gin *gin.Context) {
	var request model.UserRequest
	decoder := json.NewDecoder(gin.Request.Body)
	decoder.Decode(&request)

	gin.JSON(http.StatusOK, request)
}
