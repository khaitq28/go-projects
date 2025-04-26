package task

import (
	"fmt"
	"todo-app/internal/model"
	"todo-app/internal/user"
)

type TaskService interface {
	DeleteById(id uint) error
	Toggle(id uint) error
	Create(userId uint, title string, description string) error
}

type taskService struct {
	taskRepository TaskRepository
	userRepository user.UserRepository
}

func NewTaskService(taskRepo TaskRepository, userRepo user.UserRepository) TaskService {
	service := taskService{taskRepository: taskRepo, userRepository: userRepo}
	return &service
}

func (service *taskService) Create(userId uint, title string, description string) error {
	usr, err := service.userRepository.GetById(userId)
	if usr == nil {
		fmt.Println("Error : ", err)
		return err
	}
	task := model.Task{UserID: userId, Title: title, Des: description}
	return service.taskRepository.Create(&task)
}

func (service *taskService) DeleteById(id uint) error {
	return service.taskRepository.DeleteById(id)
}

func (service *taskService) Toggle(id uint) error {

	return service.taskRepository.ToggleTask(id)
}
