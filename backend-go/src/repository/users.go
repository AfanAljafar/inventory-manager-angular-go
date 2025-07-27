package repository

import (
	"backend-go/src/models"

	"gorm.io/gorm"
)

func AuthUser(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
