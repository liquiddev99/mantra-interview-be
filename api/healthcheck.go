package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
