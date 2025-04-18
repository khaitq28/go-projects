package user

import "fmt"

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(50);not null;"`
	Email string `gorm:"type:varchar(100);not null;unique;"`
}

func (u *User) PrintOut() {
	fmt.Printf("name: %s, email: %s\n", u.Name, u.Email)
}
