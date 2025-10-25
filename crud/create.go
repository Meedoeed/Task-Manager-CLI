package crud

import (
	"fmt"
	"os"
	"task-cli/constructors"
	"task-cli/storage"
	"task-cli/task"
)

func Create(file *os.File, desc string) ([]task.Task, error) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	tasklist, err := storage.LoadData(file)
	if err != nil {
		return nil, fmt.Errorf("error in loadfile")
	}
	t, err := constructors.NewTask(file, desc)
	if err != nil {
		return nil, fmt.Errorf("error in task creation")
	}
	tasklist = append(tasklist, t)
	return tasklist, nil
}
