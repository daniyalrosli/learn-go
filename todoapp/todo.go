package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	task        string
	isCompleted bool
}

func main() {
	todoList := []Todo{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Todo List App ---")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark task as completed")
		fmt.Println("4. Remove task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			scanner.Scan()
			task := scanner.Text()
			todoList = append(todoList, Todo{task: task, isCompleted: false})
			fmt.Println("Task added successfully!")
		case "2":
			listTasks(todoList)
		case "3":
			listTasks(todoList)
			fmt.Print("Enter task number to mark as completed: ")
			scanner.Scan()
			index := getIndex(scanner.Text(), len(todoList))
			if index != -1 {
				todoList[index].isCompleted = true
				fmt.Println("Task marked as completed!")
			}
		case "4":
			listTasks(todoList)
			fmt.Print("Enter task number to remove: ")
			scanner.Scan()
			index := getIndex(scanner.Text(), len(todoList))
			if index != -1 {
				todoList = append(todoList[:index], todoList[index+1:]...)
				fmt.Println("Task removed successfully!")
			}
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func listTasks(todoList []Todo) {
	if len(todoList) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}
	for i, todo := range todoList {
		status := " "
		if todo.isCompleted {
			status = "X"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, todo.task)
	}
}

func getIndex(input string, max int) int {
	var index int
	_, err := fmt.Sscanf(strings.TrimSpace(input), "%d", &index)
	if err != nil || index < 1 || index > max {
		fmt.Println("Invalid task number.")
		return -1
	}
	return index - 1
}
