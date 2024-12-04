package main

import (
	"golang-crud-api/controller"
	"golang-crud-api/repository"
	"golang-crud-api/service"
	_ "golang-crud-api/utils/docs"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task Management API
// @version 1.0
// @description This is a simple CRUD API for qubic ball assessment

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()

	taskRepository := &repository.TaskRepository{}

	taskService := service.NewTaskService(taskRepository)

	taskController := controller.NewTaskController(taskService)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("api/v1")

	// Task Group
	taskGroup := api.Group("/tasks")
	{
		taskGroup.GET("/", taskController.GetAllTasks)
		taskGroup.POST("/", taskController.CreateTask)
	}

	// Start the server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
