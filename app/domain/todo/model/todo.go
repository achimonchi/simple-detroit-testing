package model

import "time"

type Todo struct {
	Id          int // auto increment
	Name        string
	Description string
	CreatedAt   time.Time
}
