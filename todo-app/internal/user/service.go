package user

import "todo_app/internal/model"

type UserService interface {
	CreateUser(name, email string) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepository UserRepository
}

func (userService *userService) CreateUser(name, email string) (*model.User, error) {
	user := &model.User{Name: name, Email: email}
	err := userService.userRepository.Create(user)
	return user, err
}

func (userService *userService) GetUserById(id uint) (*model.User, error) {
	return userService.userRepository.GetById(id)
}

func (userService *userService) GetAllUsers() ([]*model.User, error) {
	return userService.userRepository.GetAll()
}

func (userService *userService) DeleteUser(id uint) error {
	return userService.userRepository.DeleteById(id)
}

func NewUserService(userRepository UserRepository) UserService {
	userService := userService{userRepository: userRepository}
	return &userService
}
