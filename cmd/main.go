package main

import (
	"os"

	"github.com/KaitoHasei/go-todo-api/internal/controllers"
	"github.com/KaitoHasei/go-todo-api/internal/models"
	"github.com/KaitoHasei/go-todo-api/internal/repositories"
	"github.com/KaitoHasei/go-todo-api/internal/services"
	"github.com/KaitoHasei/go-todo-api/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	db.AutoMigrate(&models.Todo{})

	repo := repositories.NewTodoRepository(db)
	service := services.NewTodoService(repo)
	controller := controllers.NewTodoController(service)

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/todos", controller.GetAll)
		v1.GET("/todos/:id", controller.GetById)
		v1.POST("/todos", controller.Create)
		v1.PUT("/todos/:id", controller.Update)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
