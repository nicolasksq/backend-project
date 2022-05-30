package server

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"server/api"
	"server/server/responses"
)

const (
	HOST = "localhost"
	PORT = ":8081"

	defaultEndpoint       = "/"
	peopleGetEndpoint     = "/people/"
	peopleGetByIDEndpoint = "/people/:id"

	phoneNumberField = "phone_number"
	firstnameField   = "first_name"
	lastnameField    = "last_name"
)

type Server struct {
	PeopleAPI api.PeopleGetter
}

func NewServer(peopleAPI api.PeopleGetter) Server {
	return Server{
		PeopleAPI: peopleAPI,
	}
}

// Respond to GET /people with a 200 OK response containing all people in the system
// Respond to GET /people/:id with a 200 OK response containing the requested person or a 404 Not Found response if the :id doesn't exist
// Respond to GET /people?first_name=:first_name&last_name=:last_name with a 200 OK response containing the people with that first and last name or an empty array if no people were found
// Respond to GET /people?phone_number=:phone_number with a 200 OK response containing the people with that phone number or an empty array if no people were found
func (s Server) Start() error {
	router := gin.Default()

	router.GET(defaultEndpoint, func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"hey": "hope you're having a nice day. :)"})
	})

	router.GET(peopleGetEndpoint, s.Get)
	router.GET(peopleGetByIDEndpoint, s.GetByID)

	if err := router.Run(HOST + PORT); err != nil {
		return err
	}
	return nil
}

func (s Server) Get(c *gin.Context) {
	firstname, lastname, phonenumber := c.Query(firstnameField), c.Query(lastnameField), c.Query(phoneNumberField)
	noParams := firstname == "" && lastname == "" && phonenumber == ""
	// if there is no params in url we get all.
	if noParams {
		res := s.PeopleAPI.GetAll()
		response := responses.PeopleResponse{
			People: res,
			Total:  len(res),
		}
		c.IndentedJSON(http.StatusOK, response)
		return
	}

	// if there is params in url we try to get from there.
	res, err := s.PeopleAPI.GetByFields(firstname, lastname, phonenumber)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "you should provide all params needed"})
		return
	}
	response := responses.PeopleResponse{
		People: res,
		Total:  len(res),
	}
	c.IndentedJSON(http.StatusOK, response)
}

func (s Server) GetByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.PeopleAPI.GetByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, responses.PersonResponse{Person: res})
}
