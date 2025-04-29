package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"todo-app/internal/appcontext"
)

var context, _ = appcontext.NewAppContext()

func main() {
	fmt.Println("Hello CLI")
	for {
		showHelp()
		var choice string
		fmt.Print("\nEnter your choice: ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error: Please enter a valid number!")
			continue
		}

		switch choice {
		case "1":
			createNewTask()
		case "2":
			listAllTasks()
		case "5":
			addANewUser()
		case "6":
			listAllUsers()
		case "7":
			deleteAnUser()
		case "C", "c":
			clearScreen()
		case "q", "Q":
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Print("\nPress Enter to continue...")
		fmt.Scanln()
	}

}

func createNewTask() {
	taskService := context.TaskService
	var (
		userId uint
		title  = ""
		desc   = ""
	)
	fmt.Print("\nEnter User ID: ")
	fmt.Scanln(&userId)
	fmt.Print("\nEnter Task title: ")
	fmt.Scanln(&title)
	fmt.Print("\nEnter Task description: ")
	fmt.Scanln(&desc)
	eror := taskService.Create(userId, title, desc)
	if eror != nil {
		fmt.Println("Error: ", eror)
	} else {
		fmt.Println("Task created successfully!")
	}
}

func listAllTasks() {
	taskService := context.TaskService
	tasks, err := taskService.GetAllTasks()
	if err == nil {
		for _, task := range tasks {
			task.PrintOut()
		}
	} else {
		fmt.Println("Error: ", err)
	}
}

func deleteAnUser() {
	var userId uint
	fmt.Print("\nEnter userId to delete ")
	fmt.Scanln(&userId)
	if error := context.UserService.DeleteUser(userId); error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("User deleted successfully!")
	}
}

func addANewUser() {
	userService := context.UserService
	name, email := "", ""
	fmt.Print("\nEnter your name: ")
	fmt.Scanln(&name)
	fmt.Print("\nEnter your email: ")
	fmt.Scanln(&email)

	_, eror := userService.CreateUser(name, email)
	if eror != nil {
		fmt.Println("Error: ", eror)
	} else {
		fmt.Println("User created successfully!")
	}
}

func listAllUsers() {
	userService := context.UserService
	users, _ := userService.GetAllUsers()
	for _, user := range users {
		user.PrintOut()
	}
}

func showHelp() {
	fmt.Println("Usage: todo-cli ")
	fmt.Println("Type your choice:")
	fmt.Println("   1: 		Add a new task")
	fmt.Println("   2: 		List all tasks")
	fmt.Println("   3: 		Mark a task as complete")
	fmt.Println("   4: 		Delete a task")
	fmt.Println("----------------------------------")
	fmt.Println("   5: 		Add a new user")
	fmt.Println("   6: 		List all user")
	fmt.Println("   7: 		Delete an user")
	fmt.Println(" C/c: 		Clean screen")
	fmt.Println(" q/Q: 		Exit")

}

func clearScreen() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		fmt.Println("Windows is supported.")
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
