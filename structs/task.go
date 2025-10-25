package structs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"desc"`
	Status      string `json:"status"`
	Created_at  string `json:"created"`
	Updated_at  string `json:"upd"`
}

func LoadData() (tasks []Task, err error) {
	file, err := os.OpenFile("save.json", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("error in file opening")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var task Task
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

func SaveData(t Task) error {
	file, err := os.OpenFile("save.json", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("error in save func: %s", err)
	}
	defer file.Close()
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("error in save func: %s", err)
	}
	data = append(data, '\n')
	_, err = file.WriteString(string(data))
	if err != nil {
		return fmt.Errorf("error in save func: %s", err)
	}
	return nil
}

func SaveAllData(tasklist []Task) error {
	for _, v := range tasklist {
		err := SaveData(v)
		if err != nil {
			return fmt.Errorf("can't save data")
		}
	}
	return nil
}
