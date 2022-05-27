package db_test

import (
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
