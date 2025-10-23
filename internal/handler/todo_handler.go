package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func (h *TodoHandler) HandleDeleteTodo(ctx *gin.Context) {
	strId := ctx.Param("id")

	todoId, err := strconv.Atoi(strId)
	if err != nil {
		util.CtxErrResponse(
			ctx,
			http.StatusBadRequest,
			"ID param should an integer",
			err,
		)
		return
	}

	if err := h.repo.SoftDeleteTodo(ctx, todoId); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			util.CtxErrResponse(
				ctx,
				http.StatusNotFound,
				fmt.Sprintf("No matching todo w/ ID, %d", todoId),
				err,
			)
			return
		}

		util.CtxErrResponse(
			ctx,
			http.StatusInternalServerError,
			"Server error to delete todo",
			err,
		)
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success:    true,
		Message:    fmt.Sprintf("Succesfully delete todo w/ ID, %d", todoId),
		StatusCode: http.StatusOK,
	})
}

func (h *TodoHandler) HandleUpdateTodo(ctx *gin.Context) {
	strId := ctx.Param("id")

	todoId, err := strconv.Atoi(strId)
	if err != nil {
		util.CtxErrResponse(
			ctx,
			http.StatusBadRequest,
			"ID param should an integer",
			err,
		)
		return
	}

	var editBody model.EditTodoBody
	if err := ctx.ShouldBindJSON(&editBody); err != nil {
		util.CtxErrResponse(
			ctx,
			http.StatusBadRequest,
			"Invalid request body",
			err,
		)
		return
	}

	updateData := make(map[string]any)
	if editBody.Title != nil {
		updateData["title"] = *editBody.Title
	}
	if editBody.Description != nil {
		updateData["description"] = *editBody.Description
	}
	if editBody.Completed != nil {
		updateData["completed"] = *editBody.Completed
	}

	updatedTodo, err := h.repo.UpdateTodo(ctx, todoId, updateData)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			util.CtxErrResponse(
				ctx,
				http.StatusNotFound,
				fmt.Sprintf("No matching todo w/ ID, %d", todoId),
				err,
			)
			return
		}

		util.CtxErrResponse(
			ctx,
			http.StatusInternalServerError,
			"Server error to update todo",
			err,
		)
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Success:    true,
		Data:       updatedTodo,
		Message:    fmt.Sprintf("Succesfully update todo w/ ID, %d", todoId),
		StatusCode: http.StatusOK,
	})
}
