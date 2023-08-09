package main

import (
	"github.com/AthanatiusC/mandiri-miniproject/user-service/cmd/db"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/config"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/controller"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/internal/service"
	"github.com/AthanatiusC/mandiri-miniproject/user-service/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	db, err := db.InitDB(cfg.DatabaseConfig, cfg.SecretConfig.Database)
	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	controller := controller.NewController(cfg, service)

	router := gin.Default()
	router.Use(controller.JWTMiddleware)
	router.GET("/users", controller.GetUsers)
	router.PUT("/user/:id", controller.UpdateUser)
	router.POST("/user", controller.CreateUser)
	router.DELETE("/user/:id", controller.DeleteUser)

	router.Run(":8080")
}
