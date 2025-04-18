package user

type UserService interface {
	CreateUser(name, email string) (*User, error)
	GetUserById(id uint) (*User, error)
	GetAllUsers() ([]*User, error)
}

type userService struct {
	userRepository UserRepository
}

func (userService *userService) CreateUser(name, email string) (*User, error) {
	user := &User{Name: name, Email: email}
	err := userService.userRepository.Create(user)
	return user, err
}

func (userService *userService) GetUserById(id uint) (*User, error) {
	return userService.userRepository.GetById(id)
}

func (userService *userService) GetAllUsers() ([]*User, error) {
	return userService.userRepository.GetAll()
}

func NewUserService(userRepository UserRepository) UserService {
	userService := userService{userRepository: userRepository}
	return &userService
}
