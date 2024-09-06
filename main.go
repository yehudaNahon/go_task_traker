package main

import (
	"fmt"
	"os"
	"strconv"
)

func handleHelp() {
	fmt.Println("Usage: task <command>")
	fmt.Println("Commands:")
	fmt.Println("  add <description> - Add a new task")
	fmt.Println("  delete <id> - Delete a task")
	fmt.Println("  update <id> <description> - Update a task")
	fmt.Println("  list - List all tasks")
	fmt.Println("  mark-in-progress <id> - Mark a task as in progress")
	fmt.Println("  mark-done <id> - Mark a task as done")
}

func handleAdd(task_list TaskList, args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide a description")
		return
	}
	task := task_list.Add(args[1])
	fmt.Printf("Task added with ID %d\n", task.ID)
}

func handleDelete(task_list TaskList, args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide an ID")
		return
	}

	// Convert string to int
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	task_list.Delete(id)
}

func handleUpdate(task_list TaskList, args []string) {
	if len(args) < 3 {
		fmt.Println("Please provide an ID and description")
		return
	}

	// Convert string to int
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	task_list.Update(id, args[2])
}

func handleList(task_list TaskList) {
	tasks := task_list.List()
	for _, task := range tasks {
		fmt.Printf("%d: %s - %s\n", task.ID, task.Description, task.Status)
	}
}

func handleMarkInProgress(task_list TaskList, args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide an ID")
		return
	}

	// Convert string to int
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	task_list.MarkInProgress(id)
}

func handleMarkDone(task_list TaskList, args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide an ID")
		return
	}

	// Convert string to int
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	task_list.MarkDone(id)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No command provided")
		return
	}

	task_list := Load("tasks.json")

	command := args[0]
	switch command {
	case "add":
		handleAdd(task_list, args)
	case "delete":
		handleDelete(task_list, args)
	case "update":
		handleUpdate(task_list, args)
	case "list":
		handleList(task_list)
	case "mark-in-progress":
		handleMarkInProgress(task_list, args)
	case "mark-done":
		handleMarkDone(task_list, args)
	default:
		handleHelp()
		return
	}

	err := task_list.Save("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks")
	}

}
