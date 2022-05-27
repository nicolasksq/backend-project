package transforms

import (
	"server/api/entities"
	"server/db/models"
)

func DbPeopleToPeople(personsDB []*models.Person) []*entities.Person {
	var people = make([]*entities.Person, len(personsDB))
	for i := range personsDB {
		people[i] = &entities.Person{
			ID:          personsDB[i].ID,
			FirstName:   personsDB[i].FirstName,
			LastName:    personsDB[i].LastName,
			PhoneNumber: personsDB[i].PhoneNumber,
		}
	}
	return people
}

func DbPersonToPerson(personsDB *models.Person) *entities.Person {
	var p *entities.Person
	if personsDB != nil {
		p = &entities.Person{
			ID:          personsDB.ID,
			FirstName:   personsDB.FirstName,
			LastName:    personsDB.LastName,
			PhoneNumber: personsDB.PhoneNumber,
		}
	}
	return p
}
