package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/cqs-repository-pattarn/usecase"
)

func NewUpdateTask(uc *usecase.UpdateTaskUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Println(err)
			return
		}

		input := struct {
			Name     string `json:"name"`
			Priority string `json:"priority"`
			Status   string `json:"status"`
		}{}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			log.Println(err)
			return
		}

		result := uc.Exec(ctx, &usecase.UpdateTaskUsecaseDTO{
			ID:       id,
			Name:     input.Name,
			Priority: input.Priority,
			Status:   input.Status,
		})
		if result.Err != nil {
			log.Println(result.Err)
			return
		}

		ctx.SecureJSON(http.StatusNoContent, nil)
	}
}
