package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"task-cli/task"
)

func SaveData(file *os.File, t task.Task) error {
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

func SaveAllData(file *os.File, tasklist []task.Task) error {
	for _, v := range tasklist {
		err := SaveData(file, v)
		if err != nil {
			return fmt.Errorf("can't save data")
		}
	}
	return nil
}
