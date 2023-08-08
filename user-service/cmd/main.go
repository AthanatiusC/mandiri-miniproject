package main

import (
	"github.com/AthanatiusC/mandiri-miniproject/user-service/controller"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/internal/service"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	repository := repository.NewRepository()
	service := service.NewService(repository)
	controller := controller.NewController(service)

	router := gin.Default()

	router.GET("/users", controller.GetUsers)
	router.PUT("/user/:id", controller.UpdateUser)
	router.POST("/user/:id", controller.CreateUser)
	router.DELETE("/user/:id", controller.DeleteUser)

	router.Run(":8080")
}
