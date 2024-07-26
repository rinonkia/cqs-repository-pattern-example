package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/cqs-repository-pattarn/usecase"
)

func NewGetTask(uc *usecase.GetTaskUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err)
			return
		}

		result := uc.Exec(ctx, &usecase.GetTaskUsecaseDTO{ID: id})
		if result.Err != nil {
			log.Println(result.Err)
			return
		}

		ctx.SecureJSON(http.StatusOK, gin.H{
			"task": result.Data,
		})
	}
}
