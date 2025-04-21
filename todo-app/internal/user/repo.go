package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id uint) (*User, error)
	DeleteById(id uint) error
	Create(user *User) error
	GetAll() ([]*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	r := repository{db: db}
	return &r
}
func (r *repository) GetAll() ([]*User, error) {
	var users []*User
	if err := r.db.Preload("Tasks").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) DeleteById(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *repository) GetById(id uint) (*User, error) {
	var user User
	if err := r.db.Preload("Tasks").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
