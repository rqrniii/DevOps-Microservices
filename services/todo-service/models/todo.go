package models

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
