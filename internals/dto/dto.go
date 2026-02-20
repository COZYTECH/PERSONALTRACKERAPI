package dto

type CreateWorkoutRequest struct {
	Name     string `json:"name" binding:"required"`
	Duration int    `json:"duration" binding:"required"`
	Calories int    `json:"calories"`
}

type WorkoutResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	Calories int    `json:"calories"`
}
