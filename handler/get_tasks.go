package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/cqs-repository-pattarn/usecase"
)

func NewGetTasks(uc *usecase.GetTasksUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := uc.Exec(ctx, &usecase.GetTasksUsecaseDTO{
			Status: ctx.Param("status"),
		})
		if result.Err != nil {
			log.Println(result.Err)
			return
		}

		ctx.SecureJSON(http.StatusOK, gin.H{
			"tasks": result.Data,
		})
	}
}
