package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello CLI")
	showHelp()
}

func showHelp() {
	fmt.Println("Usage: todo-cli ")
	fmt.Println("Type your choice:")
	fmt.Println("	1: Add a new task")
	fmt.Println("   2: List all tasks")
	fmt.Println("   3: Mark a task as complete")
	fmt.Println("   4: Delete a task")
	fmt.Println("   5: Clean screen")
	fmt.Println("   q/Q: Exit")

}
