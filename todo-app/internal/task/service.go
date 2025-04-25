package task

type TaskService interface {
	DeleteById(id uint) error
	Toggle(id uint) error
	Create(task *Task) error
}

type taskService struct {
	taskRepository TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	service := taskService{taskRepository: repo}
	return &service
}

func (service *taskService) Create(task *Task) error {
	return service.taskRepository.Create(task)
}

func (service *taskService) DeleteById(id uint) error {
	return service.taskRepository.DeleteById(id)
}

func (service *taskService) Toggle(id uint) error {
	return service.taskRepository.ToggleTask(id)
}
