package main

import (
	"log"

	"gantt-task-view-project/backend/internal/api"
	"gantt-task-view-project/backend/internal/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	err := db.Connect("root", "n61224n61224", "task_manager")
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	router.GET("/api/tasks", api.GetAllTasks)
	router.POST("/api/tasks", api.CreateTask)
	router.GET("/api/tasks/:id", api.GetTaskByID)
	router.PUT("/api/tasks/:id", api.UpdateTask)
	router.DELETE("/api/tasks/:id", api.DeleteTask)

	log.Println("Server running on :8080")
	router.Run(":8080")
}
