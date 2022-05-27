package models_test

import (
	"fmt"
	"testing"

	"server/db/models"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestModelsPerson(t *testing.T) {
	uuid := uuid.NewV1()
	expected := fmt.Sprintf("{\"id\":\"%s\",\"first_name\":\"Nicolas\",\"last_name\":\"Andreoli\",\"phone_number\":\"1130876000\"}", uuid.String())
	p := models.Person{
		ID:          uuid,
		FirstName:   "Nicolas",
		LastName:    "Andreoli",
		PhoneNumber: "1130876000",
	}
	result, err := p.ToJSON()
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, result)
}
