package router

import (
	"github.com/gin-gonic/gin"
	"github.com/metgag/itrix-challenge/internal/handler"
	"github.com/metgag/itrix-challenge/internal/repository"
	"gorm.io/gorm"
)

func InitTodoRoute(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewTodoRepo(db)
	handler := handler.NewTodoHandler(repo)

	todoGroup := r.Group("api/v1/todos")
	{
		todoGroup.POST("", handler.HandleCreatePost)
		todoGroup.DELETE("/:id", handler.HandleDeleteTodo)
		todoGroup.PATCH("/:id", handler.HandleUpdateTodo)
	}
}
