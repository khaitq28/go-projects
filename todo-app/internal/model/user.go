package model

import "fmt"

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(50);not null;"`
	Email string `gorm:"type:varchar(100);not null;unique;"`
	Tasks []Task `gorm:"foreignKey:UserID"`
}

func (u *User) PrintOut() {
	fmt.Println("===================\n")
	fmt.Printf("User Id: %d,  name: %s, email: %s\n", u.ID, u.Name, u.Email)
	fmt.Println("Tasks:\n")
	for _, task := range u.Tasks {
		task.PrintOut()
	}
	fmt.Println("===================\n")
}
