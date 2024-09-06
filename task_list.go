package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Status string

const (
	Todo       Status = "TODO"
	InProgress Status = "IN_PROGRESS"
	Done       Status = "DONE"
)

type Task struct {
	ID                   int
	Description          string
	Status               Status
	CreatedAt, UpdatedAt time.Time
}

type TaskListInterface interface {
	Add(description string) Task
	Delete(id int)
	Update(id int, description string) Task
	List() []Task
	MarkInProgress(id int) Task
	MarkDone(id int)
}

type TaskList struct {
	Tasks map[int]Task
}

func Load(filename string) TaskList {
	// Load tasks from file
	list := TaskList{
		Tasks: make(map[int]Task),
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File Not Found, creating new task list")
		return list
	}

	err = json.Unmarshal(data, &list)
	if err != nil {
		fmt.Println("Invalid JSON data")
		return list
	}

	return list

}

func (j TaskList) Save(filename string) error {
	data, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data")
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing file")
		return err
	}

	return nil
}

func (j *TaskList) Add(description string) Task {
	task := Task{
		ID:          len(j.Tasks) + 1,
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	j.Tasks[task.ID] = task
	return task
}

func (j *TaskList) Delete(id int) {
	delete(j.Tasks, id)
}

func (j *TaskList) Update(id int, description string) Task {
	task := j.Tasks[id]
	task.Description = description
	task.UpdatedAt = time.Now()
	j.Tasks[id] = task
	return task
}

func (j TaskList) List() []Task {
	tasks := []Task{}
	for _, task := range j.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (j *TaskList) MarkInProgress(id int) Task {
	task := j.Tasks[id]
	task.Status = InProgress
	task.UpdatedAt = time.Now()
	j.Tasks[id] = task
	return task
}

func (j *TaskList) MarkDone(id int) {
	task := j.Tasks[id]
	task.Status = Done
	task.UpdatedAt = time.Now()
	j.Tasks[id] = task
}
