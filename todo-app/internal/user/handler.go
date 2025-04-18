package user

import (
	"net/http"
	"strconv"
)
import "github.com/gin-gonic/gin"

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	handler := UserHandler{userService: userService}
	return &handler
}

func (h *UserHandler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/users")
	r.POST("/", h.createUser)
	r.GET("/:id", h.getUserByID)
	r.GET("/", h.welcome)
}

func (h *UserHandler) getUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := h.userService.GetUserById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) createUser(c *gin.Context) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user, err := h.userService.CreateUser(req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world, welcome to the user service!")
}
