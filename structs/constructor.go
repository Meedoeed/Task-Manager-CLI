package structs

import (
	"fmt"
	"slices"
	"time"
)

func GenerateID() (id int, err error) {
	tasks, err := LoadData()
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

func NewTask(description string) (Task, error) {
	TaskId, err := GenerateID()
	if err != nil && TaskId == -1 {
		return Task{}, fmt.Errorf("err in Id creation")
	}
	t := Task{
		Id:          TaskId,
		Description: description,
		Status:      "todo",
		Created_at:  time.Now().Format("15:04:05"),
		Updated_at:  time.Now().Format("15:04:05"),
	}
	if description != "" {
		err := SaveData(t)
		if err != nil {
			return t, fmt.Errorf("save error")
		}
		return t, nil
	} else {
		return Task{}, (fmt.Errorf("description can't be empty"))
	}
}
