package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rinonkia/cqs-repository-pattarn/handler"
	"github.com/rinonkia/cqs-repository-pattarn/repository/command"
	"github.com/rinonkia/cqs-repository-pattarn/repository/query"
	"github.com/rinonkia/cqs-repository-pattarn/repository/record"
	"github.com/rinonkia/cqs-repository-pattarn/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type envConfig struct {
	host       string
	port       string
	dbUser     string
	dbPassword string
	dbName     string
}

func main() {
	c, err := loadEnv()
	if err != nil {
		log.Fatal(fmt.Errorf("load env error. %w", err))
	}

	db, err := dbSetup(c)
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

func dbSetup(c *envConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		c.dbUser, c.dbPassword, c.host, c.port, c.dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&record.Task{}); err != nil {
		return nil, err
	}

	return db, nil
}

func loadEnv() (*envConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, fmt.Errorf("DB_HOST cannot be an empty string")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return nil, fmt.Errorf("DB_PORT cannot be an empty string")
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		return nil, fmt.Errorf("MYSQL_USER cannot be an empty string")
	}

	mysqlUserPassword := os.Getenv("MYSQL_USER_PASSWORD")
	if mysqlUserPassword == "" {
		return nil, fmt.Errorf("MYSQL_USER_PASSWORD cannot be an empty string")
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("DATABASE_NAME cannot be an empty string")
	}

	return &envConfig{
		host:       host,
		port:       port,
		dbUser:     mysqlUser,
		dbPassword: mysqlUserPassword,
		dbName:     dbName,
	}, nil
}
