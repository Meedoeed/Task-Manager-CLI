package task

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"desc"`
	Status      string `json:"status"`
	Created_at  string `json:"created"`
	Updated_at  string `json:"upd"`
}
