package model

import (
	"fmt"
	"time"
)

type Task struct {
	ID         uint      `gorm:"primaryKey" json:"ID"`
	UserID     uint      `gorm:"not null" json:"UserID"`
	Title      string    `gorm:"type:varchar(50);not null" json:"title"`
	Des        string    `gorm:"type:varchar(100);" json:"des"`
	Status     string    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	FinishedAt time.Time `gorm:"type:datetime;default:null" json:"finishedAt"`
}

func (t Task) PrintOut() {
	fmt.Printf("Task ID: %d, User ID: %d,   Title: %s,  Des: %s, Status: %s\n", t.ID, t.UserID, t.Title, t.Des, t.Status)
}
