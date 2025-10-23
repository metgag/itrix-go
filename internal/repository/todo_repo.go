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

func (r *TodoRepo) SoftDeleteTodo(
	ctx context.Context, todoID int,
) error {
	todo := model.Todo{}

	return r.db.WithContext(ctx).Delete(&todo, todoID).Error
}

func (r *TodoRepo) UpdateTodo(
	ctx context.Context, todoID int, body map[string]any,
) (model.Todo, error) {
	todo := model.Todo{}
	if err := r.db.WithContext(ctx).
		Model(&todo).
		Where("id = ?", todoID).
		Updates(body).
		Error; err != nil {
		return model.Todo{}, err
	}

	if err := r.db.WithContext(ctx).
		First(&todo, todoID).
		Error; err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}
