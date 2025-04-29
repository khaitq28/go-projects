package task

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"todo_app/internal/model"
)

type TaskRepository interface {
	DeleteById(id uint) error
	Create(task *model.Task) error
	ToggleTask(id uint) error
	FindById(id uint) (*model.Task, error)
	FindByUserId(userId uint) ([]*model.Task, error)
	FindAll() ([]*model.Task, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) TaskRepository {
	repo := repository{db: db}
	return &repo
}

func (rep *repository) FindAll() ([]*model.Task, error) {
	var tasks []*model.Task
	err := rep.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (rep *repository) FindByUserId(userId uint) ([]*model.Task, error) {
	var tasks []*model.Task
	err := rep.db.Where("UserID = ?", userId).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repo *repository) FindById(id uint) (*model.Task, error) {
	var task model.Task
	err := repo.db.Preload("Tasks").First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Not found")
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}
func (repo *repository) DeleteById(id uint) error {
	return repo.db.Delete(&model.Task{}, id).Error
}

func (repo *repository) Create(task *model.Task) error {
	return repo.db.Create(task).Error
}

func (repo *repository) ToggleTask(id uint) error {
	var task model.Task
	err := repo.db.Preload("Tasks").First(&task, id).Error
	if err != nil {
		fmt.Println("Not found")
		return err
	}
	if task.Status == "pending" {
		task.Status = "finished"
	} else {
		task.Status = "pending"
	}
	return nil
}
