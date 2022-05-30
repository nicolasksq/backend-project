package server_test

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"server/api/entities"
	"server/api/mocks"
	"server/server"
	"server/server/responses"
	"testing"
)

type PeopleServerTestSuite struct {
	suite.Suite
	mockApi *mocks.PeopleGetter
}

func TestPeopleServerSuite(t *testing.T) {
	suite.Run(t, new(PeopleServerTestSuite))
}

func (p *PeopleServerTestSuite) mockedServer() server.Server {
	return server.NewServer(p.mockApi)
}

func (p *PeopleServerTestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)
	p.mockApi = new(mocks.PeopleGetter)
}

func (p *PeopleServerTestSuite) TestGetByGetAll() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	s := p.mockedServer()
	id := uuid.NewV1()

	//expected
	expectedResponse := responses.PeopleResponse{
		People: []*entities.Person{{
			ID:          id,
			FirstName:   "nick",
			LastName:    "andreoli",
			PhoneNumber: "123456",
		}},
		Total: 1,
	}

	expectedCode := http.StatusOK

	p.mockApi.On("GetAll").Return([]*entities.Person{{
		ID:          id,
		FirstName:   "nick",
		LastName:    "andreoli",
		PhoneNumber: "123456",
	}}, nil).Once()

	s.Get(c)

	b, _ := ioutil.ReadAll(w.Body)
	var actualBody responses.PeopleResponse
	_ = json.Unmarshal(b, &actualBody)

	//asserts
	p.Equal(expectedCode, w.Code)
	p.Equal(expectedResponse, actualBody)
	p.mockApi.AssertExpectations(p.T())
}

func (p *PeopleServerTestSuite) TestGetByPhoneNumber() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	s := p.mockedServer()
	id := uuid.NewV1()

	//expected
	expectedResponse := responses.PeopleResponse{
		People: []*entities.Person{{
			ID:          id,
			FirstName:   "nick",
			LastName:    "andreoli",
			PhoneNumber: "123456",
		}},
		Total: 1,
	}
	expectedCode := http.StatusOK

	p.mockApi.On("GetByFields", "", "", "123456").Return([]*entities.Person{{
		ID:          id,
		FirstName:   "nick",
		LastName:    "andreoli",
		PhoneNumber: "123456",
	}}, nil).Once()

	req, _ := http.NewRequest("GET", "/people?phone_number=123456", w.Body)
	c.Request = req
	s.Get(c)

	var actualBody responses.PeopleResponse
	b, _ := ioutil.ReadAll(w.Body)
	_ = json.Unmarshal(b, &actualBody)

	//asserts
	p.Equal(expectedCode, w.Code)
	p.Equal(expectedResponse, actualBody)
	p.mockApi.AssertExpectations(p.T())
}

func (p *PeopleServerTestSuite) TestGetByFirstnameAndLastName() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	s := p.mockedServer()
	id := uuid.NewV1()

	//expected
	expectedResponse := responses.PeopleResponse{
		People: []*entities.Person{{
			ID:          id,
			FirstName:   "nick",
			LastName:    "andreoli",
			PhoneNumber: "123456",
		}},
		Total: 1,
	}
	expectedCode := http.StatusOK

	p.mockApi.On("GetByFields", "nick", "andreoli", "").Return([]*entities.Person{{
		ID:          id,
		FirstName:   "nick",
		LastName:    "andreoli",
		PhoneNumber: "123456",
	}}, nil).Once()

	req, _ := http.NewRequest("GET", "/people?first_name=nick&last_name=andreoli", w.Body)
	c.Request = req
	s.Get(c)

	var actualBody responses.PeopleResponse
	b, _ := ioutil.ReadAll(w.Body)
	_ = json.Unmarshal(b, &actualBody)

	//asserts
	p.Equal(expectedCode, w.Code)
	p.Equal(expectedResponse, actualBody)
	p.mockApi.AssertExpectations(p.T())
}

func (p *PeopleServerTestSuite) TestGetByFirstnameAndLastNameErrMissingFirstname() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	s := p.mockedServer()

	expectedCode := http.StatusBadRequest

	p.mockApi.On("GetByFields", "", "andreoli", "").Return(nil, errors.New("need a valid param to search")).Once()

	req, _ := http.NewRequest("GET", "/people?last_name=andreoli", w.Body)
	c.Request = req
	s.Get(c)

	//asserts
	p.Equal(expectedCode, w.Code)
	p.mockApi.AssertExpectations(p.T())
}

func (p *PeopleServerTestSuite) TestGetByID() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	s := p.mockedServer()
	id := uuid.FromStringOrNil("81eb745b-3aae-400b-959f-748fcafafd81")

	//expected
	expectedResponse := responses.PersonResponse{
		Person: &entities.Person{
			ID:          id,
			FirstName:   "nick",
			LastName:    "andreoli",
			PhoneNumber: "123456",
		}}
	expectedCode := http.StatusOK

	p.mockApi.On("GetByID", id).Return(&entities.Person{
		ID:          id,
		FirstName:   "nick",
		LastName:    "andreoli",
		PhoneNumber: "123456",
	}, nil).Once()

	req, _ := http.NewRequest("GET", "/people/"+id.String(), nil)
	c.Params = gin.Params{gin.Param{
		Key:   "id",
		Value: id.String(),
	}}
	c.Request = req

	s.GetByID(c)

	var actualBody responses.PersonResponse
	b, _ := ioutil.ReadAll(w.Body)
	_ = json.Unmarshal(b, &actualBody)

	//asserts
	p.Equal(expectedCode, w.Code)
	p.Equal(expectedResponse, actualBody)
	p.mockApi.AssertExpectations(p.T())
}

func (p *PeopleServerTestSuite) TestGetByBadRequest() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	s := p.mockedServer()
	id := uuid.FromStringOrNil("")

	//expected
	expectedCode := http.StatusBadRequest

	req, _ := http.NewRequest("GET", "/people/"+id.String(), nil)
	c.Request = req

	s.GetByID(c)

	var actualBody responses.PersonResponse
	b, _ := ioutil.ReadAll(w.Body)
	_ = json.Unmarshal(b, &actualBody)

	//asserts
	p.Equal(expectedCode, w.Code)
	p.mockApi.AssertExpectations(p.T())
}

func (p *PeopleServerTestSuite) TestGetByIDNotFound() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	s := p.mockedServer()
	id := uuid.FromStringOrNil("81eb745b-3aae-400b-959f-748fcafafd81")

	//expected
	expectedCode := http.StatusNotFound

	p.mockApi.On("GetByID", id).Return(nil, errors.New("user ID Not found")).Once()

	req, _ := http.NewRequest("GET", "/people/"+id.String(), nil)
	c.Params = gin.Params{gin.Param{
		Key:   "id",
		Value: id.String(),
	}}
	c.Request = req

	s.GetByID(c)

	var actualBody responses.PersonResponse
	b, _ := ioutil.ReadAll(w.Body)
	_ = json.Unmarshal(b, &actualBody)

	//asserts
	p.Equal(expectedCode, w.Code)
	p.mockApi.AssertExpectations(p.T())
}
