package db

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go-crud/src/domain"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

// RDS dynamodbStore implements domain.PetDB with an mongo storage
type RDS struct {
	Client *sqlx.DB
}

func NewRDS() (domain.PetDB, error) {
	db, err := sqlx.Connect("mysql", "root:password@pet.c1g47jkuu5ob.us-west-2.rds.amazonaws.com:3306/pet")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	return RDS{Client: db}, err
}

func (ms RDS) Get(id uuid.UUID) (*domain.Pet, error) {
	var pet *domain.Pet
	err := ms.Client.Get(&pet, "SELECT * FROM pets WHERE id=$1", id.String())
	if err != nil {
		return nil, err
	}

	return pet, nil
}

func (ms RDS) List() ([]*domain.Pet, error) {
	var pets []*domain.Pet
	err := ms.Client.Select(&pets, "SELECT * FROM pets")
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func (ms RDS) Create(pet *domain.Pet) error {
	_, err := ms.Client.NamedExec(`INSERT INTO pets (id,category,age,price) VALUES (:id,:category,:age,:price)`, &pet)
	return err
}

func (ms RDS) Delete(id uuid.UUID) error {
	_, err := ms.Client.Exec("DELETE FROM pets WHERE id=$1", id.String())
	return err
}
