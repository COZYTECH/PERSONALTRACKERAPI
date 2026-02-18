package handlers

import (
	"net/http"
	"strconv"

	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/services"

	"github.com/gin-gonic/gin"
)

// WorkoutRequest struct for binding JSON in requests
type WorkoutRequest struct {
	Name     string `json:"name" binding:"required"`
	Duration int    `json:"duration" binding:"required"` // in minutes
	Calories int    `json:"calories"`
}

// CreateWorkout - POST /workouts
func CreateWorkout(c *gin.Context) {
	userID := c.GetInt("user_id")

	var req WorkoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// err := services.CreateWorkout(userID, req.Name, req.Duration, req.Calories)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusCreated, gin.H{"message": "Workout created"})
		err := services.CreateWorkout(userID, req.Name, req.Duration, req.Calories)
	if err != nil {
		if err.Error() == "workout already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Workout created"})
}

// GetWorkouts - GET /workouts
func GetWorkouts(c *gin.Context) {
	userID := c.GetInt("user_id")

	workouts, err := services.GetWorkouts(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workouts)
}

// UpdateWorkout - PUT /workouts/:id
func UpdateWorkout(c *gin.Context) {
	userID := c.GetInt("user_id")
	workoutID, _ := strconv.Atoi(c.Param("id"))

	var req WorkoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateWorkout(userID, workoutID, req.Name, req.Duration, req.Calories)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workout updated"})
}

// DeleteWorkout - DELETE /workouts/:id
func DeleteWorkout(c *gin.Context) {
	userID := c.GetInt("user_id")
	workoutID, _ := strconv.Atoi(c.Param("id"))

	err := services.DeleteWorkout(userID, workoutID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workout deleted"})
}
