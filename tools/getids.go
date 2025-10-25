package tools

import "task-cli/structs"

func GetIds(tasklist []structs.Task) []int {
	var idlist []int
	for _, v := range tasklist {
		idlist = append(idlist, v.Id)
	}
	return idlist
}
