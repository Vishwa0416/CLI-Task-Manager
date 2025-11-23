package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Task struct represents a task
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// Global task list
var tasks []Task
var dataFile = "tasks.json"

// Load tasks from JSON file
func loadTasks() {
	file, err := ioutil.ReadFile(dataFile)
	if err != nil {
		// File might not exist yet
		tasks = []Task{}
		return
	}
	json.Unmarshal(file, &tasks)
}

// Save tasks to JSON file
func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	ioutil.WriteFile(dataFile, data, 0644)
}

// Add a new task
func addTask(description string) {
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	newTask := Task{ID: id, Description: description, Done: false}
	tasks = append(tasks, newTask)
	saveTasks()
	fmt.Println("Task added:", description)
}

// List all tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("Tasks:")
	for _, t := range tasks {
		status := "❌"
		if t.Done {
			status = "✅"
		}
		fmt.Printf("%d. [%s] %s\n", t.ID, status, t.Description)
	}
}

// Mark a task as done
func markDone(id int) {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			saveTasks()
			fmt.Println("Task marked as done:", t.Description)
			return
		}
	}
	fmt.Println("Task ID not found.")
}

// Delete a task
func deleteTask(id int) {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Println("Task deleted:", t.Description)
			return
		}
	}
	fmt.Println("Task ID not found.")
}

// CLI menu
func main() {
	loadTasks()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Task Manager ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter task description: ")
			desc, _ := reader.ReadString('\n')
			addTask(strings.TrimSpace(desc))
		case "2":
			listTasks()
		case "3":
			fmt.Print("Enter task ID to mark done: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			markDone(id)
		case "4":
			fmt.Print("Enter task ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			deleteTask(id)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
