package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo_app/internal/appcontext"
	"todo_app/internal/task"
	"todo_app/internal/user"

	_ "todo_app/cmd/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My App API
// @version 1.0
// @description This is a sample Gin application with Swagger
// @host localhost:8080
// @BasePath /
func main() {

	fmt.Println("Hello World")

	ctx, err := appcontext.NewAppContext()
	if err != nil {
		fmt.Println("Failed to init context:", err)
		return
	}

	userHandler := user.NewUserHandler(ctx.UserService)
	taskHandler := task.NewTaskHandler(ctx.TaskService)

	r := gin.Default()
	api := r.Group("/api/v1")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userHandler.RegisterRoutes(api)
	taskHandler.RegisterRoutes(api)

	r.Run(":8080")
}
