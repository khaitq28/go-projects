package user

import (
	"gorm.io/gorm"
	"todo_app/internal/model"
)

type UserRepository interface {
	GetById(id uint) (*model.User, error)
	DeleteById(id uint) error
	Create(user *model.User) error
	GetAll() ([]*model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	r := repository{db: db}
	return &r
}
func (r *repository) GetAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Preload("Tasks").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *repository) DeleteById(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *repository) GetById(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.Preload("Tasks").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
