package api_test

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"server/api"
	"server/api/transforms"
	"server/db/mocks"
	"server/db/models"
	"testing"
)

type PeopleApiTestSuite struct {
	suite.Suite
	mockDb    *mocks.FinderPeople
	apiPeople api.PeopleGetter
}

func TestPeopleApiSuite(t *testing.T) {
	suite.Run(t, new(PeopleApiTestSuite))
}

func (p *PeopleApiTestSuite) SetupTest() {
	p.mockDb = new(mocks.FinderPeople)
	p.apiPeople = api.NewPeopleAPI(p.mockDb)
}

func (p *PeopleApiTestSuite) TestGetAllPeopleSuccess() {
	peopleDB := []*models.Person{
		{
			ID:          uuid.Must(uuid.FromString("81eb745b-3aae-400b-959f-748fcafafd81")),
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "+1 (800) 555-1212",
		}}
	expected := transforms.DbPeopleToPeople(peopleDB)

	p.mockDb.On("GetAllPeople").Return(peopleDB).Once()
	result := p.apiPeople.GetAll()

	p.mockDb.AssertExpectations(p.T())
	assert.Equal(p.T(), expected, result)
}

func (p *PeopleApiTestSuite) TestGetByIDSuccess() {
	uuid := uuid.FromStringOrNil("81eb745b-3aae-400b-959f-748fcafafd81")
	personDB := &models.Person{
		ID:          uuid,
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "+1 (800) 555-1212",
	}
	expected := transforms.DbPersonToPerson(personDB)

	p.mockDb.On("FindPersonByID", uuid).Return(personDB, nil).Once()
	result, err := p.apiPeople.GetByID(uuid)

	p.mockDb.AssertExpectations(p.T())
	assert.Nil(p.T(), err)
	assert.Equal(p.T(), expected, result)
}

func (p *PeopleApiTestSuite) TestGetByIDErr() {
	id := uuid.FromStringOrNil("81eb745b-3aae-400b-959f-748fcafafd81")
	expected := errors.New("not found")

	p.mockDb.On("FindPersonByID", id).Return(nil, expected).Once()
	result, err := p.apiPeople.GetByID(id)

	p.mockDb.AssertExpectations(p.T())
	assert.Nil(p.T(), result)
	assert.Equal(p.T(), expected, err)
}

func (p *PeopleApiTestSuite) TestGetByPhoneNumberSuccess() {
	phone := "+1 (800) 555-1212"
	peopleDB := []*models.Person{{
		ID:          uuid.FromStringOrNil("81eb745b-3aae-400b-959f-748fcafafd81"),
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: phone,
	}}
	expected := transforms.DbPeopleToPeople(peopleDB)

	p.mockDb.On("FindPeopleByPhoneNumber", phone).Return(peopleDB, nil).Once()
	result, err := p.apiPeople.GetByFields("", "", phone)

	p.mockDb.AssertExpectations(p.T())
	assert.Nil(p.T(), err)
	assert.Equal(p.T(), expected, result)
}

func (p *PeopleApiTestSuite) TestGetByNamesSuccess() {
	firstname := "John"
	lastname := "Doe"
	peopleDB := []*models.Person{{
		ID:          uuid.FromStringOrNil("81eb745b-3aae-400b-959f-748fcafafd81"),
		FirstName:   firstname,
		LastName:    lastname,
		PhoneNumber: "+1 (800) 555-1212",
	}}
	expected := transforms.DbPeopleToPeople(peopleDB)

	p.mockDb.On("FindPeopleByNames", firstname, lastname).Return(peopleDB, nil).Once()
	result, err := p.apiPeople.GetByFields(firstname, lastname, "")

	p.mockDb.AssertExpectations(p.T())
	assert.Nil(p.T(), err)
	assert.Equal(p.T(), expected, result)
}

func (p *PeopleApiTestSuite) TestGetByFieldsInvalidFields() {
	expected := errors.New("need a valid param to search")
	result, err := p.apiPeople.GetByFields("", "", "")

	assert.Nil(p.T(), result)
	assert.Equal(p.T(), expected, err)
}
