package mocks

import (
	uuid "github.com/satori/go.uuid"
	"server/db/models"
)

// People is the data source for the People RESTful service.
var People = []*models.Person{
	{
		ID:          uuid.Must(uuid.FromString("81eb745b-3aae-400b-959f-748fcafafd81")),
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "+1 (800) 555-1212",
	},
	{
		ID:          uuid.Must(uuid.FromString("5b81b629-9026-450d-8e46-da4f8c7bd513")),
		FirstName:   "Jane",
		LastName:    "Doe",
		PhoneNumber: "+1 (800) 555-1313",
	},
	{
		ID:          uuid.Must(uuid.FromString("df12ce76-767b-4bf0-bccb-816745df9e70")),
		FirstName:   "Brian",
		LastName:    "Smith",
		PhoneNumber: "+44 7700 900077",
	},
	// This is another person with the name John Doe
	{
		ID:          uuid.Must(uuid.FromString("135af595-aa86-4bb5-a8f7-df17e6148e63")),
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "+1 (800) 555-1414",
	},
	// This is another person with the phone number +44 7700 900077
	{
		ID:          uuid.Must(uuid.FromString("000ebe58-b659-422b-ab48-a0d0d40bd8f9")),
		FirstName:   "Jenny",
		LastName:    "Smith",
		PhoneNumber: "+44 7700 900077",
	},
}
