package tools

import "task-cli/task"

func GetIds(tasklist []task.Task) []int {
	var idlist []int
	for _, v := range tasklist {
		idlist = append(idlist, v.Id)
	}
	return idlist
}
