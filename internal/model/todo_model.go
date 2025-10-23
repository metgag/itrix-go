package model

import (
	"time"

	"gorm.io/gorm"
)

type TodoBody struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type Todo struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string         `gorm:"type:varchar(95);not null" json:"title"`
	Description string         `gorm:"type:text;not null" json:"description"`
	Completed   bool           `gorm:"type:boolean;default:false" json:"completed"`
	Categories  []Category     `gorm:"many2many:todo_categories"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type EditTodoBody struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

type Category struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:varchar(75);not null;unique" json:"name"`
	Todos []Todo `gorm:"many2many:todo_categories"`
}
