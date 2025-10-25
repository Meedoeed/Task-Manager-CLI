package crud

import (
	"fmt"
	"os"
	"task-cli/structs"
	"task-cli/tools"
)

const (
	EditDesc     = 1
	ChangeStatus = 2
)

func UpdateTask(id int, choice int, data any) ([]structs.Task, error) {
	tasklist, err := structs.LoadData()
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
	if id == -1 {
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
	_, err = os.OpenFile("save.json", os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("file cant be open")
	}
	err = structs.SaveAllData(tasklist)
	if err != nil {
		return nil, fmt.Errorf("save all data error")
	}
	return tasklist, nil
}
