package entity

import "time"

// Todo ...
type Todo struct {
	ID        int
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
