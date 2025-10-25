package task

import (
	"fmt"
	"time"
)

func (t *Task) EditDesc(desc string) error {
	if desc != "" {
		t.Description = desc
		t.Updated_at = time.Now().Format("15:04:05")
		return nil
	} else {
		return fmt.Errorf("the description of the task being modified cannot be empty")
	}
}

func (t *Task) ChangeStatus(status_input string) error {
	if t.Status == "completed" {
		return fmt.Errorf("task status is already 'completed'")
	} else {
		switch status_input {
		case "todo":
			t.Status = "todo"
		case "in-progress":
			t.Status = "in-progress"
		case "completed":
			t.Status = "completed"
		default:
			t.Status = "todo"
		}
		t.Updated_at = time.Now().Format("15:04:05")
		return nil
	}
}
