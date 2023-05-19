package domain

import (
	"github.com/google/uuid"
)

type Pet struct {
	ID       string  `json:"id" db:"id" `
	Category string  `json:"category" db:"category"`
	Age      int     `json:"age" db:"age"`
	Price    float64 `json:"price" db:"price"`
}

type PetSvc interface {
	Get(id uuid.UUID) (*Pet, error)
	List() ([]*Pet, error)
	Create(p *Pet) error
	Delete(id uuid.UUID) error
}

type PetDB interface {
	Get(id uuid.UUID) (*Pet, error)
	List() ([]*Pet, error)
	Create(p *Pet) error
	Delete(id uuid.UUID) error
}
