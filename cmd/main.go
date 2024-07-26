package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/cqs-repository-pattarn/handler"
	"github.com/rinonkia/cqs-repository-pattarn/repository/command"
	"github.com/rinonkia/cqs-repository-pattarn/repository/query"
	"github.com/rinonkia/cqs-repository-pattarn/usecase"
)

func main() {
	r := gin.New()

	// repository
	getTaskByIDQuery := query.NewGetTaskByIDQuery()
	getAllTaskQuery := query.NewGetAllTasksQuery()
	getTasksByStatusQuery := query.NewGetTasksByStatusQuery()
	putTaskCommand := command.NewPutTaskCommand()
	deleteTaskCommand := command.NewDeleteTaskCommand()

	// usecase
	getTaskUsecase := usecase.NewGetTaskUsecase(getTaskByIDQuery)
	addTaskUsecase := usecase.NewAddTaskUsecase(putTaskCommand)
	updateTaskUsecase := usecase.NewUpdateTaskUsecase(getTaskByIDQuery, putTaskCommand)
	deleteTaskUsecase := usecase.NewDeleteTaskUsecase(deleteTaskCommand)
	getTasksUsecase := usecase.NewGetTasksUsecase(getAllTaskQuery, getTasksByStatusQuery)

	// handler
	healthCheck := handler.NewHealthCheck()
	getTaskHandler := handler.NewGetTask(getTaskUsecase)
	addTaskHandler := handler.NewAddTask(addTaskUsecase)
	updateTaskHandler := handler.NewUpdateTask(updateTaskUsecase)
	deleteTaskHandler := handler.NewDeleteTask(deleteTaskUsecase)
	getTasksHandler := handler.NewGetTasks(getTasksUsecase)

	r.GET("/health", healthCheck)
	r.GET("/task/:id", getTaskHandler)
	r.POST("/task", addTaskHandler)
	r.PATCH("/task", updateTaskHandler)
	r.DELETE("/task/:id", deleteTaskHandler)
	r.GET("/tasks", getTasksHandler)

	log.Fatal(r.Run(":8080"))
}
