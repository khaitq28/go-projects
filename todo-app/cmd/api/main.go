package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"todo-app/internal/user"
)

func main() {

	fmt.Println("Hello World")

	dsn := "host=localhost user=todo_user password=todo_pass dbname=todo_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	r := gin.Default()
	api := r.Group("/api/v1")
	userHandler.RegisterRoutes(api)
	r.Run(":8080")
}
