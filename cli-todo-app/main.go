package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Description string
	Completed   bool
}

var todos []Todo

func displayTodos() {
	if len(todos) == 0 {
		fmt.Println("Your todo list is empty!")
		return
	}

	for i, todo := range todos {
		status := " "
		if todo.Completed {
			status = "X"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, todo.Description)
	}
}

func addTodo(description string) {
	todo := Todo{Description: description, Completed: false}
	todos = append(todos, todo)
	fmt.Println("Added:", description)
}

func completeTodo(index int) {
	if index < 1 || index > len(todos) {
		fmt.Println("Invalid task number.")
		return
	}
	todos[index-1].Completed = true
	fmt.Println("Marked as completed:", todos[index-1].Description)
}

func deleteTodo(index int) {
	if index < 1 || index > len(todos) {
		fmt.Println("Invalid task number.")
		return
	}
	todo := todos[index-1]
	todos = append(todos[:index-1], todos[index:]...)
	fmt.Println("Deleted:", todo.Description)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nTodo CLI - Choose an option:")
		fmt.Println("1. View Todos")
		fmt.Println("2. Add Todo")
		fmt.Println("3. Complete Todo")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Quit")

		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			displayTodos()
		case "2":
			fmt.Print("Enter task description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)
			addTodo(description)
		case "3":
			fmt.Print("Enter task number to complete: ")
			numberStr, _ := reader.ReadString('\n')
			number, err := strconv.Atoi(strings.TrimSpace(numberStr))
			if err != nil {
				fmt.Println("Please enter a valid number.")
				continue
			}
			completeTodo(number)
		case "4":
			fmt.Print("Enter task number to delete: ")
			numberStr, _ := reader.ReadString('\n')
			number, err := strconv.Atoi(strings.TrimSpace(numberStr))
			if err != nil {
				fmt.Println("Please enter a valid number.")
				continue
			}
			deleteTodo(number)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}
	}
}
