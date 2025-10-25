package crud

import (
	"fmt"
	"os"
	"task-cli/structs"
	"task-cli/tools"
)

func DelTask(id int) ([]structs.Task, error) {
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
	if index == -1 {
		return nil, fmt.Errorf("id is not in list")
	}

	tasklist = append(tasklist[:index], tasklist[index+1:]...)

	_, err = os.OpenFile("save.json", os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("file cant be open")
	}
	err = structs.SaveAllData(tasklist)
	if err != nil {
		return nil, fmt.Errorf("can't save data")
	}
	return tasklist, nil
}
