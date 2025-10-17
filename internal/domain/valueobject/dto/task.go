package dto

import "github.com/google/uuid"

type TaskResponse struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
}
