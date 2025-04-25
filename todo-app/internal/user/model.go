package user

import "fmt"
import "todo-app/internal/task"

type User struct {
	ID    uint        `gorm:"primaryKey"`
	Name  string      `gorm:"type:varchar(50);not null;"`
	Email string      `gorm:"type:varchar(100);not null;unique;"`
	Tasks []task.Task `gorm:"foreignKey:UserID"`
}

func (u *User) PrintOut() {
	fmt.Printf("id: %d,  name: %s, email: %s\n", u.ID, u.Name, u.Email)
}
