package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) pong(ctx *gin.Context) {

	ans := gin.H{
		"message": "pong",
	}
	ctx.JSON(http.StatusOK, ans)

}
