package main

import (
	"bufio"
	"fmt"
	"os"
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

func addNewTask (name string){
	newTask := Task{Name: name, Done: false}
	tasks = append(tasks, newTask)
	fmt.Println("Added new task successfully!")
}

func main () {
	var scanner = bufio.NewScanner(os.Stdin)

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
		
		fmt.Print("Insert an action number: ")
		fmt.Scanln(&action)

		divider()

		switch action {
		case "0":
			fmt.Println("Goodbye!")
			return
		case "1":
			fmt.Println("Add a new task")

			var task string
			fmt.Print("Insert the task description/name: ")
			scanner.Scan()
			task = scanner.Text()
			addNewTask(task)
		case "2":
			fmt.Println("Set a task as done - This will be implemented soon")
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}