package controller

import (
	"backend-go/src/models"
	"backend-go/src/repository"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Message struct {
	Message string `json:"message"`
}

func AuthUser(c echo.Context, db *gorm.DB) error {
	var request models.User
	// var user models.User

	// c.JSON(http.StatusOK, Message{Message: "Hello from Echo POST!"})

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
	// 	return c.JSON(http.StatusUnauthorized, echo.Map{"error": "email not found"})
	// }
	user, err := repository.AuthUser(db, request.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "User tidak ditemukan"})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid password"})
	}

	// Create JWT Token
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Email,
		"role":     user.UserCategory,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 1 day
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"token":   t,
		"role":    user.UserCategory,
		"user_id": user.ID,
		"name":    user.Name,
	})
}
