package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AthanatiusC/mandiri-miniproject/user-service/config"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/helper"
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
		helper.HandleErrorResponse(gctx, err)
		return
	}

	stringID, err := c.JWTGetClaimValue(gctx, "id")
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}
	id, _ := strconv.Atoi(stringID)

	response, err := c.Service.CreateUsers(int64(id), request)
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}

	gctx.JSON(response.Code, response)
}

func (c *Controller) DeleteUser(gctx *gin.Context) {
	requestUserID := gctx.Param("id")
	userID, err := strconv.Atoi(requestUserID)
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}

	response, err := c.Service.DeleteUsers(userID)
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}

	gctx.JSON(response.Code, response)
}
