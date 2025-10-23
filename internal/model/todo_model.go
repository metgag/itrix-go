package model

import (
	"time"

	"gorm.io/gorm"
)

type TodoBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Todo struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"type:varchar(95);not null"`
	Description string `gorm:"type:text;not null"`
	Completed   bool   `gorm:"type:boolean;default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
