package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/minhtri6179/manata/db/sqlc"
)

type CreateTaskRequest struct {
	Id          int         `json:"id" `
	Title       string      `json:"title" `
	Description pgtype.Text `json:"description" `
}

func (server *Server) createTask(ctx *gin.Context) {
	var req CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTaskParams{
		Title:       req.Title,
		Description: req.Description,
	}

	task, err := server.store.CreateTask(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (server *Server) getTask(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	task, err := server.store.GetTask(ctx, int32(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, task)
}

type UpdateTaskRequest struct {
	Title       string      `json:"title" `
	Description pgtype.Text `json:"description" `
	Status      string      `json:"status" `
}

func (server *Server) updateTask(ctx *gin.Context) {
	var req UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdateStatusParams{
		ID:          int32(id),
		Title:       req.Title,
		Description: req.Description,
	}
	err = server.store.UpdateStatus(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Update status successfully"})

}

func (server *Server) deleteTask(ctx *gin.Context) {
	var req UpdateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	DeletedStatus := "Doing"
	arg := db.UpdateStatusParams{
		ID:          int32(id),
		Title:       req.Title,
		Description: req.Description,
		Status:      &UpdateTaskRequest{&req.Status: &DeletedStatus},
	}
	err = server.store.UpdateStatus(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete successfully"})

}
