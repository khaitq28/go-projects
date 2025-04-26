package model

import "time"

type Task struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	Title      string    `gorm:"type:varchar(50);not null"`
	Des        string    `gorm:"type:varchar(100);"`
	Status     string    `gorm:"type:varchar(20);default:'pending'"`
	Finished   bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	FinishedAt *time.Time
}
