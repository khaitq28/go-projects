package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-app/internal/appcontext"
	"todo-app/internal/user"
)

func main() {

	fmt.Println("Hello World")

	ctx, err := appcontext.NewAppContext()
	if err != nil {
		fmt.Println("Failed to init context:", err)
		return
	}

	userHandler := user.NewUserHandler(ctx.UserService)

	r := gin.Default()
	api := r.Group("/api/v1")
	userHandler.RegisterRoutes(api)
	r.Run(":8080")
}
