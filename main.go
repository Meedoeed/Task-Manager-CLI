package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"task-cli/crud"
	"task-cli/structs"
)

const (
	helpText = `Task Manager - CLI Task Management System

				This application provides a CRUD interface for managing your tasks.
				All tasks are automatically saved in 'save.json' file.

				Available commands:
				-add-task [description]        - Add a new task with description
				-update-task [id] [description] - Update task description by ID
				-del-task [id]                 - Delete task by ID
				-change-status [id] [status]   - Change task status ('todo'/'in-progress'/'completed')
				-task-list                     - Display all tasks

				Examples:
				-add-task "Buy groceries"
				-update-task 1 "Buy groceries and cook dinner"
				-del-task 3
				-change-status 2 completed
				-task-list

				Hope you'll enjoy using it!`
)

func main() {
	//flags
	HelpPtr := flag.Bool("help", false, "a bool")
	addTask := flag.String("add-task", "", "a string var")
	editTask := flag.Bool("update-task", false, "a string var")
	changest := flag.Bool("change-status", false, "a string var")
	delTask := flag.Int("del-task", 0, "a string var")
	tasklist := flag.Bool("task-list", false, "a bool")

	flag.Parse()

	if *changest {
		args := flag.Args()
		var status string
		if len(args) >= 2 {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Error in type convertation:%s\n", err)
			}
			status = args[1]
			if status == "todo" || status == "completed" || status == "in-progress" {
				_, err = crud.UpdateTask(id, 2, status)
				if err != nil {
					fmt.Printf("Error in task editing: %s\n", err)
				} else {
					fmt.Println("Task status was updated successfully!")
				}
			} else {
				fmt.Println("Not valid status! try 'todo'/'in-progress/'completed'")
			}
		} else {
			os.Exit(1)
		}
	}

	if *delTask != 0 {
		_, err := crud.DelTask(*delTask)
		if err != nil {
			fmt.Printf("Error when deliting: %s\n", err)
		} else {
			fmt.Println("Task was deleted successfully!")
		}
	}

	if *editTask {
		args := flag.Args()
		var status string
		if len(args) >= 2 {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Error in type convertation:%s\n", err)
			}
			status = args[1]
			_, err = crud.UpdateTask(id, 1, status)
			if err != nil {
				fmt.Printf("Error in task editing: %s\n", err)
			} else {
				fmt.Println("Task description was updated successfully! ")
			}
		} else {
			os.Exit(1)
		}
	}

	if *HelpPtr {
		fmt.Println(helpText)
	}

	if *addTask != "" {
		_, err := crud.Create(*addTask)
		if err != nil {
			fmt.Printf("Error in task creation: %s\n", err)
		} else {
			fmt.Println("Task was created successfully! ")
		}
	}

	if *tasklist {
		tasklist, err := structs.LoadData()
		if err != nil {
			fmt.Printf("Error in task loading: %s\n", err)
		}
		args := flag.Args()
		if len(args) != 0 {

			switch args[0] {
			case "todo":
				for _, v := range tasklist {
					if v.Status == "todo" {
						fmt.Println("Todo list:")
						fmt.Println(v)
					}
				}
			case "in-progress":
				for _, v := range tasklist {
					if v.Status == "in-progress" {
						fmt.Println("In-progress list:")
						fmt.Println(v)
					}
				}
			}
		} else {
			for _, v := range tasklist {
				fmt.Println(v)
			}
		}
	}
}
