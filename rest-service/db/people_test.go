package db_test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"server/db/models"
	"testing"

	"server/db"
	"server/db/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PeopleTestSuite struct {
	suite.Suite
	peopleDB *db.PeopleDB
}

func TestPeopleSuite(t *testing.T) {
	suite.Run(t, new(PeopleTestSuite))
}

func (p *PeopleTestSuite) SetupTest() {
	p.peopleDB = new(db.PeopleDB)
}

func (p *PeopleTestSuite) TestGetAllPeopleSuccess() {
	expected := mocks.People
	people := p.peopleDB.AllPeople()

	assert.Equal(p.T(), expected, people)
}

func (p *PeopleTestSuite) TestFindPersonByIDSuccess() {
	expected := mocks.People[0]
	people, err := p.peopleDB.FindPersonByID(uuid.FromStringOrNil("81eb745b-3aae-400b-959f-748fcafafd81"))
	assert.Nil(p.T(), err)
	assert.Equal(p.T(), expected, people)
}

func (p *PeopleTestSuite) TestFindPersonByIDNotFound() {
	uuid := uuid.NewV1()
	expected := fmt.Sprintf("user ID %s not found", uuid.String())
	people, err := p.peopleDB.FindPersonByID(uuid)
	assert.Nil(p.T(), people)
	assert.Equal(p.T(), expected, err.Error())
}

func (p *PeopleTestSuite) TestFindPeopleByNamesSuccess() {
	expected := []*models.Person{mocks.People[0], mocks.People[3]}
	people := p.peopleDB.FindPeopleByNames("John", "Doe")
	assert.Equal(p.T(), expected, people)
}

func (p *PeopleTestSuite) TestFindPeopleByPhoneNumberSuccess() {
	expected := []*models.Person{mocks.People[2], mocks.People[4]}
	people := p.peopleDB.FindPeopleByPhoneNumber("+44 7700 900077")
	assert.Equal(p.T(), expected, people)
}
