package models

import (
	"encoding/json"

	"github.com/satori/go.uuid"
)

// Person defines a simple representation of a person
type Person struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
}

func (person *Person) ToJSON() (string, error) {
	marshaled, err := json.Marshal(person)
	if err != nil {
		return "", err
	}

	return string(marshaled[:]), nil
}
