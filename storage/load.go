package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"task-cli/task"
)

func GenerateID(file *os.File) (id int, err error) {
	_, err = file.Seek(0, 0)
	if err != nil {
		return -1, err
	}
	tasks, err := LoadData(file)
	if err != nil {
		return -1, err
	}
	var listID []int
	for _, v := range tasks {
		listID = append(listID, v.Id)
	}
	if len(listID) != 0 {
		return slices.Max(listID) + 1, nil
	} else {
		return 1, nil
	}
}

func LoadData(file *os.File) (tasks []task.Task, err error) {

	scanner := bufio.NewScanner(file)
	var task task.Task
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &task); err != nil {
			fmt.Printf("Warning: failed to unmarshal task: %v\n", err)
			continue
		}
		tasks = append(tasks, task)
	}
	if err := scanner.Err(); err != nil {
		return tasks, fmt.Errorf("scanner error: %w", err)
	}
	return tasks, nil
}
