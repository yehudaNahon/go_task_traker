package main

import (
	"fmt"
	"os"
	"strconv"
)

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
		if len(args) < 2 {
			fmt.Println("Please provide a description")
			return
		}
		task := task_list.Add(args[1])
		fmt.Printf("Task added with ID %d\n", task.ID)

	case "delete":
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

	case "update":
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

	case "list":
		tasks := task_list.List()
		for _, task := range tasks {
			fmt.Printf("%d: %s - %s\n", task.ID, task.Description, task.Status)
		}

	case "mark-in-progress":
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

	case "mark-done":
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
	default:
		fmt.Println("Invalid command")
	}

	err := task_list.Save("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks")
	}

}
