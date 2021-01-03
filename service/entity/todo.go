package entity

import "time"

// Todo ...
type Todo struct {
	ID        string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
