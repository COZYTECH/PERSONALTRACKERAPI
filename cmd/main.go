package main

import (
	"fmt"
	"log"

	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/database"
	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/handlers"
	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/middleware"
	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/utils"
	"github.com/gin-gonic/gin"
)


func main(){

	cfg := utils.LoadConfig()
	database.Init()

	r := gin.Default()

auth := r.Group("/auth")
{
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)
}

protected := r.Group("/workouts")
protected.Use(middleware.AuthMiddleware())
{
	protected.POST("/", handlers.CreateWorkout)
	protected.GET("/", handlers.GetWorkouts)
	protected.PUT("/:id", handlers.UpdateWorkout)
	protected.DELETE("/:id", handlers.DeleteWorkout)
}

log.Printf("Server running on port %s", cfg.Port)
	err := r.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

	fmt.Println("Server started successfully")
}

