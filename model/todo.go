package model

import "time"

type Todo struct {
	Id        string    `json:"id"`
	Task      string    `json:"task"`
	User      string    `json:"user"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
