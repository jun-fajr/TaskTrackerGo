package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "todo", "in-progress", "done"
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var taskFile = "tasks.json"

// LoadTasks loads the tasks from the JSON file, or initializes an empty list if the file does not exist.
func LoadTasks() ([]Task, error) {
	file, err := os.ReadFile(taskFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// SaveTasks saves the tasks list to the JSON file.
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

// AddTask adds a new task to the list.
func AddTask(description string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	id := len(tasks) + 1
	task := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, task)
	if err := SaveTasks(tasks); err != nil {
		return err
	}
	fmt.Printf("Task added successfully (ID: %d)\n", id)
	return nil
}

// UpdateTask updates the description of an existing task.
func UpdateTask(id int, description string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			if err := SaveTasks(tasks); err != nil {
				return err
			}
			fmt.Println("Task updated successfully")
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// DeleteTask removes a task from the list by ID.
func DeleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := SaveTasks(tasks); err != nil {
				return err
			}
			fmt.Println("Task deleted successfully")
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// MarkTask updates the status of a task.
func MarkTask(id int, status string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			if err := SaveTasks(tasks); err != nil {
				return err
			}
			fmt.Printf("Task marked as %s\n", status)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// ListTasks displays tasks with optional filtering by status.
func ListTasks(filter string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			fmt.Printf("ID: %d | Description: %s | Status: %s | CreatedAt: %s | UpdatedAt: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339))
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <description>")
			return
		}
		description := os.Args[2]
		if err := AddTask(description); err != nil {
			fmt.Println("Error adding task:", err)
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> <description>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		description := os.Args[3]
		if err := UpdateTask(id, description); err != nil {
			fmt.Println("Error updating task:", err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := DeleteTask(id); err != nil {
			fmt.Println("Error deleting task:", err)
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := MarkTask(id, "in-progress"); err != nil {
			fmt.Println("Error marking task:", err)
		}
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := MarkTask(id, "done"); err != nil {
			fmt.Println("Error marking task:", err)
		}
	case "list":
		filter := ""
		if len(os.Args) >= 3 {
			filter = os.Args[2]
		}
		if err := ListTasks(filter); err != nil {
			fmt.Println("Error listing tasks:", err)
		}
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: add, update, delete, mark-in-progress, mark-done, list")
	}
}
