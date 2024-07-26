package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHealthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "health ok")
		return
	}
}
