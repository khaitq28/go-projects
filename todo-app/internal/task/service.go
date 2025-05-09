package task

import (
	"todo_app/internal/model"
)

type TaskService interface {
	DeleteById(id uint) error
	Toggle(id uint) error
	Create(userId uint, title string, description string) (*model.Task, error)
	GetByUserId(userId uint) ([]*model.Task, error)
	GetAllTasks() ([]*model.Task, error)
	GetById(id uint) (*model.Task, error)
}

type taskService struct {
	taskRepository TaskRepository
}

func NewTaskService(taskRepo TaskRepository) TaskService {
	service := taskService{taskRepository: taskRepo}
	return &service
}

func (service *taskService) GetByUserId(userId uint) ([]*model.Task, error) {
	return service.taskRepository.FindByUserId(userId)
}

func (service *taskService) GetById(id uint) (*model.Task, error) {
	return service.taskRepository.FindById(id)
}
func (service *taskService) GetAllTasks() ([]*model.Task, error) {
	return service.taskRepository.FindAll()
}

func (service *taskService) Create(userId uint, title string, description string) (*model.Task, error) {
	task := model.Task{UserID: userId, Title: title, Des: description}
	err := service.taskRepository.Create(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (service *taskService) DeleteById(id uint) error {
	return service.taskRepository.DeleteById(id)
}

func (service *taskService) Toggle(id uint) error {

	return service.taskRepository.ToggleTask(id)
}
