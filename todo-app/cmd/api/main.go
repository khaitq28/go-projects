package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand/v2"
	"net/http"
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

	// Initialize data
	initData()

	userHandler := user.NewUserHandler(ctx.UserService)
	taskHandler := task.NewTaskHandler(ctx.TaskService)

	r := gin.New()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Max-Age", "300")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL)
		c.Next()
	})

	api := r.Group("/api/v1")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userHandler.RegisterRoutes(api)
	taskHandler.RegisterRoutes(api)

	r.Run(":8080")
}

func initData() {
	ctx, err := appcontext.NewAppContext()
	if err != nil {
		fmt.Println("Failed to init context:", err)
		return
	}
	userService := ctx.UserService
	taskService := ctx.TaskService
	users, _ := userService.GetAllUsers()
	for _, user := range users {
		for i := 0; i < 10; i++ {
			id := rand.Uint32()
			title := fmt.Sprintf("Task %d", id)
			desc := fmt.Sprintf("Description %d", id)
			taskService.Create(user.ID, title, desc)
		}
	}
}
