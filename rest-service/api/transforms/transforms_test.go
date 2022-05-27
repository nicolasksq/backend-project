package transforms_test

import (
	"testing"

	"server/api/entities"
	"server/api/transforms"
	"server/db/models"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestDbPeopleToPeople(t *testing.T) {
	var (
		id        = uuid.NewV1()
		firstName = "Nick"
		lastName  = "Andreoli"
		phone     = "1130876000"
	)

	expected := []*entities.Person{{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phone,
	}}

	result := transforms.DbPeopleToPeople([]*models.Person{{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phone,
	}})

	assert.Equal(t, expected, result)
}

func TestDbPersonToPerson(t *testing.T) {
	var (
		id        = uuid.NewV1()
		firstName = "Nick"
		lastName  = "Andreoli"
		phone     = "1130876000"
	)

	expected := &entities.Person{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phone,
	}
	result := transforms.DbPersonToPerson(&models.Person{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phone,
	})

	assert.Equal(t, expected, result)
}

func TestDbPersonToPersonWithNilDB(t *testing.T) {
	var expected *entities.Person
	result := transforms.DbPersonToPerson(nil)

	assert.Equal(t, expected, result)
}
