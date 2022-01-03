package model

type Job struct {
	ID        int64  `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}
