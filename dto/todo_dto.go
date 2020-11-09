package dto

// TodoBaseDTO ..
type TodoBaseDTO struct {
	Value string `json:"value"`
}

// TodoDTO ...
type TodoDTO struct {
	ID int `json:"id"`
	TodoBaseDTO
}
