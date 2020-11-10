package dto

import "time"

// TodoBaseDTO ..
type TodoBaseDTO struct {
	Value string `json:"value"`
}

// TodoDTO ...
type TodoDTO struct {
	ID string `json:"id"`
	TodoBaseDTO
}

// TodoDetailDTO ...
type TodoDetailDTO struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	TodoDTO
}
