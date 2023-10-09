package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/pgtype"
	db "github.com/minhtri6179/manata/db/sqlc"
)

type createUserRequest struct {
	Username    string           `json:"username" binding:"required,alphanum"`
	Password    string           `json:"password" binding:"required,min=6"`
	FirstName   string           `json:"first_name" binding:"required"`
	LastName    string           `json:"last_name" binding:"required"`
	DateOfBirth pgtype.Timestamp `json:"date_of_birth" binding:"required"`
	Email       string           `json:"email" binding:"required,email"`
}
type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	arg := db.CreateUserParams{
		UserName:       req.Username,
		HashedPassword: hashPassword,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		DateOfBirth:    req.DateOfBirth,
		Email:          req.Email,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}
