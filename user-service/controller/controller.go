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
	err := gctx.ShouldBindJSON(&request)
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}

	stringLevel, err := c.JWTGetClaimValue(gctx, "level")
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}
	accessLevel, _ := strconv.Atoi(stringLevel)

	response, err := c.Service.GetUsers(accessLevel, request)
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}

	gctx.JSON(http.StatusOK, response)
}

func (c *Controller) UpdateUser(gctx *gin.Context) {
	var request model.UserRequest
	decoder := json.NewDecoder(gctx.Request.Body)
	decoder.Decode(&request)

	stringLevel, err := c.JWTGetClaimValue(gctx, "level")
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}
	accessLevel, _ := strconv.Atoi(stringLevel)

	requestUserID := gctx.Param("id")
	userID, err := strconv.Atoi(requestUserID)
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}

	response, err := c.Service.UpdateUsers(accessLevel, int64(userID), request)
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}

	gctx.JSON(http.StatusOK, response)
}

func (c *Controller) CreateUser(gctx *gin.Context) {
	var request model.UserRequest
	err := gctx.ShouldBindJSON(&request)
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}

	stringLevel, err := c.JWTGetClaimValue(gctx, "level")
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}
	accessLevel, _ := strconv.Atoi(stringLevel)

	response, err := c.Service.CreateUsers(accessLevel, request)
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

	stringLevel, err := c.JWTGetClaimValue(gctx, "levl")
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}
	accessLevel, _ := strconv.Atoi(stringLevel)

	response, err := c.Service.DeleteUsers(accessLevel, userID)
	if err != nil {
		helper.HandleErrorResponse(gctx, err)
		return
	}

	gctx.JSON(response.Code, response)
}
