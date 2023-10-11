package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/minhtri6179/manata/common"
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
	nullStatus := db.NullStatus{
		Status: db.StatusDoing,
		Valid:  true,
	}
	arg := db.CreateTaskParams{
		Title:       req.Title,
		Description: req.Description,
		Status:      nullStatus,
	}

	task, err := server.store.CreateTask(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrCannotCreateEntity(task.Title, err))
		return
	}

	ctx.JSON(http.StatusOK, task.ID)
}

func (server *Server) getTask(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	task, err := server.store.GetTask(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrEntityNotFound(task.Title, err))
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
		ctx.JSON(http.StatusBadRequest, common.ErrCannotCreateEntity(req.Title, err))
		return
	}

	status, err := StringToStatus(req.Status)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, common.ErrCannotCreateEntity(req.Title, err))
		return
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Status: %s\n", status)
	}
	nullStatus := db.NullStatus{
		Status: status,
		Valid:  true,
	}

	arg := db.UpdateStatusParams{
		ID:          int32(id),
		Title:       req.Title,
		Description: req.Description,
		Status:      nullStatus,
	}
	err = server.store.UpdateStatus(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrCannotUpdateEntity(req.Title, err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Update status successfully"})

}
func (server *Server) deleteTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	task, err := server.store.GetTask(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	args := db.UpdateStatusParams{
		ID:          int32(id),
		Title:       task.Title,
		Description: task.Description,
		Status:      db.NullStatus{Status: db.StatusDeleted, Valid: true},
	}

	err = server.store.UpdateStatus(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrCannotDeleteEntity(task.Title, err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete task successfully"})
}

// Maybe handle paging for request params in the future

func (server *Server) listTask(ctx *gin.Context) {
	var paging common.Pagging
	if err := ctx.ShouldBind(&paging); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	paging.HandlePaging()
	arg := db.ListTasksParams{
		Limit:  paging.Limit,
		Offset: ((paging.Page - 1) * paging.Limit),
	}

	tasks, err := server.store.ListTasks(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrCannotListEntity("Can not show", err))
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
func StringToStatus(str string) (db.Status, error) {
	switch str {
	case "Doing":
		return db.StatusDoing, nil
	case "Done":
		return db.StatusDone, nil
	case "Inprocess":
		return db.StatusInprocess, nil
	case "Deleted":
		return db.StatusDeleted, nil
	default:
		return "", fmt.Errorf("invalid Status value: %s", str)
	}
}
