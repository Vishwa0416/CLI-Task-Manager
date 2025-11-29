package tasks

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type Task struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
    Done        bool   `json:"done"`
}

var tasks []Task
var dataFile = "tasks.json"

func LoadTasks() {
    file, err := ioutil.ReadFile(dataFile)
    if err != nil {
        tasks = []Task{}
        return
    }
    json.Unmarshal(file, &tasks)
}

func SaveTasks() {
    data, _ := json.MarshalIndent(tasks, "", "  ")
    ioutil.WriteFile(dataFile, data, 0644)
}

func AddTask(description string) {
    id := 1
    if len(tasks) > 0 {
        id = tasks[len(tasks)-1].ID + 1
    }
    newTask := Task{ID: id, Description: description, Done: false}
    tasks = append(tasks, newTask)
    SaveTasks()
    fmt.Println("Task added:", description)
}

func ListTasks() {
    LoadTasks()
    if len(tasks) == 0 {
        fmt.Println("No tasks found.")
        return
    }
    for _, t := range tasks {
        status := "❌"
        if t.Done {
            status = "✅"
        }
        fmt.Printf("%d. [%s] %s\n", t.ID, status, t.Description)
    }
}

func MarkDone(id int) {
    LoadTasks()
    for i, t := range tasks {
        if t.ID == id {
            tasks[i].Done = true
            SaveTasks()
            fmt.Println("Marked done:", t.Description)
            return
        }
    }
    fmt.Println("Task ID not found.")
}

func DeleteTask(id int) {
    LoadTasks()
    for i, t := range tasks {
        if t.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            SaveTasks()
            fmt.Println("Task deleted:", t.Description)
            return
        }
    }
    fmt.Println("Task ID not found.")
}
