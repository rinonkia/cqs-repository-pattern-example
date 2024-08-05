package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/cqs-repository-pattarn/usecase"
)

func NewAddTask(uc *usecase.AddTaskUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := struct {
			Name     string `json:"name"`
			Priority string `json:"priority"`
		}{}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			log.Println(err)
			return
		}

		err := uc.Exec(ctx, &usecase.AddTaskUsecaseDTO{
			Name:     input.Name,
			Priority: input.Priority,
		})
		if err != nil {
			log.Println(err)
			return
		}

		ctx.SecureJSON(http.StatusNoContent, nil)
	}
}
