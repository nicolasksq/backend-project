package main

import (
	"fmt"
	"server/api"
	"server/db"
	"server/server"
)

//Web services are our bread and butter. Our services talk to each other over gRPC and REST.
// The rest-service directory contains a simple Person model and a set of sample data that needs a REST service in front of it.
// This service should:
//
// Respond with JSON output
//
// You can implement the service with go's built-in routines or import a framework or router if you like. The Person model and all of the backend code is in the rest-service/pkg/models/person.go file, the service should be initialized in rest-service/main.go, and should run by running go run main.go from the rest-service directory.
//
// Implementing the service is a good start, but are there any extras you can throw in? How would you test this service? How would you audit it? How would an ops person audit it?
func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	s := server.NewServer(api.NewPeopleAPI(db.PeopleDB{}))
	if err := s.Start(); err != nil {
		panic(err)
	}
}
