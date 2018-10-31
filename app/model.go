package app

type Task struct {
	Id    int16  `json:"id"`
	Title string `json:"title"`
	Done  string `json:"boolean"`
}
