package main

import (
	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/handlers"
	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/middleware"

	"github.com/gin-gonic/gin"
)


func main(){
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

r.Run(":8080")

}