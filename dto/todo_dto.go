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

// PageableDto base dto to for pagenable response
type PageableDto struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Count      int `json:"count"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// TodoPageableDto Dto to return pageable to Todo
type TodoPageableDto struct {
	TodoDTO `json:"records"`
	PageableDto
}
