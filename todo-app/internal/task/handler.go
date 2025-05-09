package task

import (
	"net/http"
	"strconv"
)
import "github.com/gin-gonic/gin"

type TaskHandler struct {
	taskService TaskService
}

func NewTaskHandler(taskService TaskService) *TaskHandler {
	handler := TaskHandler{taskService: taskService}
	return &handler
}

func (h *TaskHandler) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/tasks")
	r.POST("", h.createTask)
	r.GET("", h.getAllTasks)
	r.GET("/:id", h.getTaskById)
	//r.GET("/", h.welcome)
}

type RequestBody struct {
	UserId uint   `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Desc   string `json:"description" binding:"required"`
}

func (h *TaskHandler) createTask(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	task, err := h.taskService.Create(body.UserId, body.Title, body.Desc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create task"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) getAllTasks(c *gin.Context) {
	tasks, err := h.taskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get all tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello world, welcome to the Tasks service!")
}

func (h *TaskHandler) getTaskById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := h.taskService.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found : " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
