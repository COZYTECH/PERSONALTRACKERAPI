package services

import (
	"errors"
	"time"

	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/database"
	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/models"
)

// CreateWorkout creates a workout for a user
func CreateWorkout(userID int, name string, duration int, calories int) error {
	if name == "" {
		return errors.New("workout name is required")
	}
	if duration <= 0 {
		return errors.New("duration must be greater than zero")
	}

	_, err := database.DB.Exec(
		"INSERT INTO workouts (user_id, name, duration, calories, created_at) VALUES (?, ?, ?, ?, ?)",
		userID, name, duration, calories, time.Now(),
	)
	return err
}

// GetWorkouts retrieves workouts for a user
func GetWorkouts(userID int) ([]models.Workout, error) {
	rows, err := database.DB.Query(
		"SELECT id, name, duration, calories, created_at FROM workouts WHERE user_id = ?",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workouts := []models.Workout{}
	for rows.Next() {
		var w models.Workout
		err := rows.Scan(&w.ID, &w.Name, &w.Duration, &w.Calories, &w.CreatedAt)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, w)
	}

	return workouts, nil
}

// UpdateWorkout updates a user's workout
func UpdateWorkout(userID int, workoutID int, name string, duration int, calories int) error {
	if name == "" {
		return errors.New("workout name is required")
	}
	if duration <= 0 {
		return errors.New("duration must be greater than zero")
	}

	_, err := database.DB.Exec(
		"UPDATE workouts SET name=?, duration=?, calories=? WHERE id=? AND user_id=?",
		name, duration, calories, workoutID, userID,
	)
	return err
}

// DeleteWorkout deletes a user's workout
func DeleteWorkout(userID int, workoutID int) error {
	_, err := database.DB.Exec(
		"DELETE FROM workouts WHERE id=? AND user_id=?",
		workoutID, userID,
	)
	return err
}
