package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Status string `json:"status"`
	// ... other fields
}
