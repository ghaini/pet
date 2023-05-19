package app

import (
	"github.com/google/uuid"
	"go-crud/src/domain"
)

// petSvc implements domain.PetSvc
type petSvc struct {
	DB domain.PetDB
}

func NewPetSvc(db domain.PetDB) domain.PetSvc {
	return petSvc{
		DB: db,
	}
}

func (ps petSvc) Get(id uuid.UUID) (*domain.Pet, error) {
	return ps.DB.Get(id)

}

func (ps petSvc) List() ([]*domain.Pet, error) {
	return ps.DB.List()
}

func (ps petSvc) Create(pet *domain.Pet) error {
	pet.ID = uuid.New().String()
	return ps.DB.Create(pet)
}

func (ps petSvc) Delete(id uuid.UUID) error {
	return ps.DB.Delete(id)
}
