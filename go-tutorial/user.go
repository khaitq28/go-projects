package main

import (
	"go-tutorial/internal/maths"
	"go-tutorial/internal/stringutils"
)

type User struct {
	name  string
	age   int
	email string
}

func (user *User) Print() {
	user.age = maths.Mul(2, user.age)
	println("name:", user.name)
	println("age:", user.age)
	println("email:", user.email)
}

func (user *User) IsEqual(user2 *User) bool {
	return user.name == user2.name &&
		user.age == user2.age &&
		user.email == user2.email
}

func filter(users []User, filterFunc func(User) bool) []User {
	var result []User
	for _, user := range users {
		if filterFunc(user) {
			result = append(result, user)
		}
	}
	return result
}

func filterAdult(users []User) []User {
	return filter(users, isAdult)
}

func filterTeenager(users []User) []User {
	return filter(users, isTeenager)
}

func filterByName(users []User, name string) []User {
	return filter(users, func(user User) bool {
		return stringutils.FindString(user.name, name)
	})
}

func isAdult(user User) bool {
	return user.age >= 18
}

func isTeenager(user User) bool {
	return user.age >= 13 && user.age < 18
}
