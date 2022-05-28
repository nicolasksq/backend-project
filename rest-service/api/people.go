package api

import (
	"errors"

	"server/api/entities"
	"server/api/transforms"
	"server/db"

	uuid "github.com/satori/go.uuid"
)

type PeopleGetter interface {
	GetAll() []*entities.Person
	GetByID(id uuid.UUID) (*entities.Person, error)
	GetByFields(firstname, lastname, phoneNumber string) ([]*entities.Person, error)
}

type People struct {
	peopleDB db.FinderPeople
}

func NewPeopleAPI(db db.FinderPeople) PeopleGetter {
	return People{
		peopleDB: db,
	}
}

func (p People) GetAll() []*entities.Person {
	return transforms.DbPeopleToPeople(p.peopleDB.GetAllPeople())
}

func (p People) GetByID(uuid uuid.UUID) (*entities.Person, error) {
	person, err := p.peopleDB.FindPersonByID(uuid)
	if err != nil {
		return nil, err
	}
	return transforms.DbPersonToPerson(person), nil
}

func (p People) GetByFields(firstname, lastname, phoneNumber string) ([]*entities.Person, error) {
	if firstname != "" && lastname != "" {
		peopleDB := p.peopleDB.FindPeopleByNames(firstname, lastname)
		return transforms.DbPeopleToPeople(peopleDB), nil
	}
	if phoneNumber != "" {
		peopleDB := p.peopleDB.FindPeopleByPhoneNumber(phoneNumber)
		return transforms.DbPeopleToPeople(peopleDB), nil
	}

	return nil, errors.New("need a valid param to search")
}
