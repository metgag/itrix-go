package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/metgag/itrix-challenge/internal/model"
	"github.com/metgag/itrix-challenge/internal/repository"
	"github.com/metgag/itrix-challenge/internal/util"
)

type TodoHandler struct {
	repo *repository.TodoRepo
}

func NewTodoHandler(repo *repository.TodoRepo) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (h *TodoHandler) HandleCreatePost(ctx *gin.Context) {
	var body model.TodoBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		util.CtxErrResponse(
			ctx,
			http.StatusBadRequest,
			"Invalid request body",
			err,
		)
		return
	}

	todo, err := h.repo.CreateTodo(ctx, body)
	if err != nil {
		util.CtxErrResponse(
			ctx,
			http.StatusInternalServerError,
			"Server failed to create todo",
			err,
		)
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success:    true,
		Data:       todo,
		Message:    "Todo added succesfully",
		StatusCode: http.StatusOK,
	})
}
