package db

import (
	"errors"
	"fmt"

	"server/db/mocks"
	"server/db/models"

	uuid "github.com/satori/go.uuid"
)

type FinderPeople interface {
	GetAllPeople() []*models.Person
	FindPersonByID(id uuid.UUID) (*models.Person, error)
	FindPeopleByNames(firstName, lastName string) []*models.Person
	FindPeopleByPhoneNumber(phoneNumber string) []*models.Person
}

type PeopleDB struct{}

// GetAllPeople returns all people in `people`.
func (p PeopleDB) GetAllPeople() []*models.Person {
	return mocks.People
}

// FindPersonByID searches for people in `people` the by their ID.
func (p PeopleDB) FindPersonByID(id uuid.UUID) (*models.Person, error) {
	for _, person := range mocks.People {
		if person.ID == id {
			return person, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("user ID %s not found", id.String()))
}

// FindPeopleByNames performs a case-sensitive search for people in `people` by first and last name.
func (p PeopleDB) FindPeopleByNames(firstName, lastName string) []*models.Person {
	result := make([]*models.Person, 0)
	for _, person := range mocks.People {
		if person.FirstName == firstName && person.LastName == lastName {
			result = append(result, person)
		}
	}
	return result
}

// FindPeopleByPhoneNumber searches for people in `people` by phone number.
func (p PeopleDB) FindPeopleByPhoneNumber(phoneNumber string) []*models.Person {
	result := make([]*models.Person, 0)
	for _, person := range mocks.People {
		if person.PhoneNumber == phoneNumber {
			result = append(result, person)
		}
	}
	return result
}
