package model

import (
	"fmt"
	"time"
)

type Task struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	Title      string    `gorm:"type:varchar(50);not null"`
	Des        string    `gorm:"type:varchar(100);"`
	Status     string    `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	FinishedAt time.Time `gorm:"type:datetime;default:null"`
}

func (t Task) PrintOut() {
	fmt.Printf("Task ID: %d, User ID: %d,   Title: %s,  Des: %s, Status: %s\n", t.ID, t.UserID, t.Title, t.Des, t.Status)
}
