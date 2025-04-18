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
	r.GET("/", h.getAllUsers)
	//r.GET("/", h.welcome)
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
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found : " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

type RequestBody struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (h *UserHandler) createUser(c *gin.Context) {

	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user, err := h.userService.CreateUser(body.Name, body.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) getAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get all users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world, welcome to the user service!")
}
