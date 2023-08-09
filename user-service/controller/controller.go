package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/config"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/internal/service"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/model"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Config  *config.Config
	Service *service.Service
}

func NewController(config *config.Config, service *service.Service) *Controller {
	return &Controller{
		Config:  config,
		Service: service,
	}
}

func (c *Controller) GetUsers(gctx *gin.Context) {
	var request model.UserRequest
	gctx.ShouldBindJSON(&request)
	if request.ID != 0 {
		gctx.JSON(http.StatusOK, request)
		return
	}

	gctx.JSON(http.StatusOK, request)
}

func (c *Controller) UpdateUser(gctx *gin.Context) {
	var request model.UserRequest
	decoder := json.NewDecoder(gctx.Request.Body)
	decoder.Decode(&request)

	gctx.JSON(http.StatusOK, request)
}

func (c *Controller) CreateUser(gctx *gin.Context) {
	var request model.UserRequest
	err := gctx.ShouldBindJSON(&request)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := c.Service.CreateUsers(request)
	if err != nil {
		fmt.Println("Error occured", err)
		gctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	gctx.JSON(http.StatusOK, response)
}

func (c *Controller) DeleteUser(gctx *gin.Context) {
	var request model.UserRequest
	decoder := json.NewDecoder(gctx.Request.Body)
	decoder.Decode(&request)

	gctx.JSON(http.StatusOK, request)
}
