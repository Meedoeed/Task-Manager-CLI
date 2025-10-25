package crud

import (
	"fmt"
	"os"
	"task-cli/storage"
	"task-cli/task"
	"task-cli/tools"
)

const (
	EditDesc     = 1
	ChangeStatus = 2
)

func UpdateTask(file *os.File, id int, choice int, data any) ([]task.Task, error) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
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

	str, ok := data.(string)
	if !ok {
		return nil, fmt.Errorf("not supported type, expected string")
	}
	switch choice {
	case EditDesc:
		err := tasklist[index].EditDesc(str)
		if err != nil {
			return nil, fmt.Errorf("error in edit task:%s", err)
		}
	case ChangeStatus:
		err := tasklist[index].ChangeStatus(str)
		if err != nil {
			return nil, fmt.Errorf("error in change status:%s", err)
		}
	}
	file.Close()
	file, err = os.OpenFile("save.json", os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("file cant be open")
	}
	defer file.Close()
	err = storage.SaveAllData(file, tasklist)
	if err != nil {
		return nil, fmt.Errorf("save all data error")
	}
	return tasklist, nil
}
