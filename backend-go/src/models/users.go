package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"` // cocok dengan SERIAL
	Name         string    `json:"name"`                 // cocok dengan VARCHAR
	Email        string    `json:"email" gorm:"unique"`  // cocok dengan UNIQUE
	Password     string    `json:"password"`             // cocok dengan VARCHAR
	UserCategory string    `json:"user_category"`        // cocok dengan VARCHAR + CHECK
	CreatedAt    time.Time `json:"created_at"`           // cocok dengan TIMESTAMP DEFAULT NOW()
	UpdatedAt    time.Time `json:"updated_at"`           // cocok dengan TIMESTAMP DEFAULT NOW()
}

func (User) TableName() string {
	return "inventory_management.users"
}
