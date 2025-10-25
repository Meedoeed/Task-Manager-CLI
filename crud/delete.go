package crud

import (
	"fmt"
	"os"
	"task-cli/storage"
	"task-cli/task"
	"task-cli/tools"
)

func DelTask(file *os.File, id int) ([]task.Task, error) {
	tasklist, err := storage.LoadData(file)
	if err != nil {
		return nil, fmt.Errorf("error in load file")
	}
	idlist := tools.GetIds(tasklist)
	var index int = -1
	for i, v := range idlist {
		if id == v {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, fmt.Errorf("id is not in list")
	}

	tasklist = append(tasklist[:index], tasklist[index+1:]...)

	file, err = os.OpenFile("save.json", os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("file cant be open")
	}
	err = storage.SaveAllData(file, tasklist)
	if err != nil {
		return nil, fmt.Errorf("can't save data")
	}
	return tasklist, nil
}
