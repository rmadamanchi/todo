package tasks

type Task struct {
	Id    int16  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
