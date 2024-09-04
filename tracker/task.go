package tracker

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

type Task struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time,omitempty"`
	Duration  string    `json:"duration,omitempty"`
}

var tasksFile = "tasks.json"

func SaveTask(task Task) error {
	var existingTasks []Task
	file, err := os.OpenFile(tasksFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&existingTasks)
	if err != nil && err != io.EOF {
		return err
	}

	existingTasks = append(existingTasks, task)

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	return encoder.Encode(existingTasks)
}

func LoadLastTask() (*Task, error) {
	file, err := os.OpenFile(tasksFile, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, nil
	}

	lastTask := tasks[len(tasks)-1]
	return &lastTask, nil
}

func ClearTasks() error {
	file, err := os.OpenFile(tasksFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		return err
	}
	return nil
}
