package constructors

import (
	"fmt"
	"os"
	"task-cli/storage"
	"task-cli/task"
	"time"
)

func NewTask(file *os.File, description string) (task.Task, error) {
	TaskId, err := storage.GenerateID(file)
	if err != nil && TaskId == -1 {
		return task.Task{}, fmt.Errorf("err in Id creation")
	}
	t := task.Task{
		Id:          TaskId,
		Description: description,
		Status:      "todo",
		Created_at:  time.Now().Format("15:04:05"),
		Updated_at:  time.Now().Format("15:04:05"),
	}
	if description != "" {
		err := storage.SaveData(file, t)
		if err != nil {
			return t, fmt.Errorf("save error")
		}
		return t, nil
	} else {
		return task.Task{}, (fmt.Errorf("description can't be empty"))
	}
}
