package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rinonkia/cqs-repository-pattarn/handler"
	"github.com/rinonkia/cqs-repository-pattarn/repository/command"
	"github.com/rinonkia/cqs-repository-pattarn/repository/query"
	"github.com/rinonkia/cqs-repository-pattarn/repository/record"
	"github.com/rinonkia/cqs-repository-pattarn/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := dbSetup()
	if err != nil {
		log.Fatal(fmt.Errorf("db setup error. %w", err))
	}

	r := gin.New()

	// repository
	getTaskByIDQuery := query.NewGetTaskByIDQuery(db)
	getAllTaskQuery := query.NewGetAllTasksQuery(db)
	getTasksByStatusQuery := query.NewGetTasksByStatusQuery(db)
	putTaskCommand := command.NewPutTaskCommand(db)
	deleteTaskCommand := command.NewDeleteTaskCommand(db)

	// usecase
	getTaskUsecase := usecase.NewGetTaskUsecase(getTaskByIDQuery)
	addTaskUsecase := usecase.NewAddTaskUsecase(putTaskCommand)
	updateTaskUsecase := usecase.NewUpdateTaskUsecase(getTaskByIDQuery, putTaskCommand)
	deleteTaskUsecase := usecase.NewDeleteTaskUsecase(getTaskByIDQuery, deleteTaskCommand)
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
	r.PATCH("/task/:id", updateTaskHandler)
	r.DELETE("/task/:id", deleteTaskHandler)
	r.GET("/tasks", getTasksHandler)

	log.Fatal(r.Run(":8080"))
}

func dbSetup() (*gorm.DB, error) {
	var (
		host   = "0.0.0.0"
		port   = "3306"
		user   = "crp"
		pw     = "crp"
		dbName = "crp"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", user, pw, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&record.Task{}); err != nil {
		return nil, err
	}

	return db, nil
}
