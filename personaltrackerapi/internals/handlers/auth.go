package handlers

import (
	"net/http"

	"github.com/COZYTECH/PERSONALTRACKERAPI/personaltrackerapi/cmd/internals/database"
	"github.com/COZYTECH/PERSONALTRACKERAPI/personaltrackerapi/cmd/internals/models"
	"github.com/COZYTECH/PERSONALTRACKERAPI/personaltrackerapi/cmd/internals/services"
	"github.com/gin-gonic/gin"
)

// Register a new user
func Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashed, err := services.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Insert into DB
	_, err = database.DB.Exec(
		"INSERT INTO users (email, password) VALUES (?, ?)",
		req.Email, hashed,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email may already exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login a user
func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := database.DB.QueryRow(
		"SELECT id, password FROM users WHERE email = ?",
		req.Email,
	).Scan(&user.ID, &user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password
	if err := services.CheckPassword(user.Password, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT
	token, err := services.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
