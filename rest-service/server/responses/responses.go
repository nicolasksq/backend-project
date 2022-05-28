package responses

import "server/api/entities"

type PeopleResponse struct {
	People []*entities.Person `json:"people"`
	Total  int                `json:"total,omitempty"`
}

type PersonResponse struct {
	Person *entities.Person `json:"person"`
}
