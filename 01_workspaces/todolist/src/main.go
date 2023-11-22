package main

import (
	"fmt"
)

var tasks []Task

type Task struct {
	Name string
	Done bool
}

func divider () {
	fmt.Println("--------------------")
}

func loadTasks () {
	if len(tasks) > 0 {
		for _, task := range tasks {
			if task.Done {
				fmt.Println("[x] - " + task.Name)
			} else {
				fmt.Println("[ ] - " + task.Name)
			}
		}
	} else {
		fmt.Println("No tasks yet!")
	}
}

func main () {
	divider()

	fmt.Println("Welcome to Todolist!")

	for {
		divider()

		loadTasks()

		divider()

		fmt.Println(
			"1 - Add a new task \n" +
			"2 - Set a task as done \n" +
			"0 - Exit")
		
		divider()
		
		var action string

		fmt.Scanln(&action)

		divider()

		switch action {
		case "0":
			fmt.Println("Goodbye!")
			return
		case "1":
			fmt.Println("Add a new task - This will be implemented soon")
		case "2":
			fmt.Println("Set a task as done - This will be implemented soon")
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}