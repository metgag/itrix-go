package repository

import (
	"context"

	"github.com/metgag/itrix-challenge/internal/model"
	"gorm.io/gorm"
)

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
	return &TodoRepo{db: db}
}

func (r *TodoRepo) CreateTodo(
	ctx context.Context, body model.TodoBody,
) (*model.Todo, error) {
	todo := model.Todo{
		Title:       body.Title,
		Description: body.Description,
	}

	if err := r.db.WithContext(ctx).Create(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}
