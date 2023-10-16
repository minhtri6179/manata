package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
	db "github.com/minhtri6179/manata/db/sqlc"
	"github.com/minhtri6179/manata/util"
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
	Username  string           `json:"username"`
	FullName  string           `json:"full_name"`
	Email     string           `json:"email"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

func (server *Server) registerUser(ctx *gin.Context) {
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
		if e, ok := err.(*pq.Error); ok {
			switch e.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, gin.H{"error": e.Message})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Message})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user_res := userResponse{
		Username:  user.UserName,
		FullName:  user.FirstName + " " + user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, user_res)

}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}
type loginUserResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := util.CheckPassword(req.Password, user.HashedPassword); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := server.tokenMaker.GenerateToken(user.UserName, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := loginUserResponse{
		Username: user.UserName,
		Token:    token,
	}
	ctx.JSON(http.StatusOK, res)
}
