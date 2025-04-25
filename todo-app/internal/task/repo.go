package task

import (
	"fmt"
	"gorm.io/gorm"
)

type TaskRepository interface {
	DeleteById(id uint) error
	Create(task *Task) error
	ToggleTask(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) TaskRepository {
	repo := repository{db: db}
	return &repo
}

func (repo *repository) DeleteById(id uint) error {
	return repo.db.Delete(&Task{}, id).Error
}

func (repo *repository) Create(task *Task) error {
	return repo.db.Create(task).Error
}

func (repo *repository) ToggleTask(id uint) error {
	var task Task
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
