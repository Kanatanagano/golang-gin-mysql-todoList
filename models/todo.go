package models

type Todo struct {
	ID        int    `json:"id"`
	Title     string `jspn:"title"`
	Completed bool   `json:"completed"`
}