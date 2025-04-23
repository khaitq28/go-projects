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
		case "6":
			listAllUsers()
		case "C", "c":
			clearScreen()
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Print("\nPress Enter to continue...")
		fmt.Scanln()
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
	fmt.Println("	1: 		Add a new task")
	fmt.Println("   2: 		List all tasks")
	fmt.Println("   3: 		Mark a task as complete")
	fmt.Println("   4: 		Delete a task")
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
