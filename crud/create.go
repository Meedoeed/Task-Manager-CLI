package crud

import (
	"fmt"
	"task-cli/structs"
)

func Create(desc string) ([]structs.Task, error) {
	tasklist, err := structs.LoadData()
	if err != nil {
		return nil, fmt.Errorf("error in loadfile")
	}
	t, err := structs.NewTask(desc)
	if err != nil {
		return nil, fmt.Errorf("error in task creation")
	}
	tasklist = append(tasklist, t)
	return tasklist, nil
}
