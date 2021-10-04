package main

import (
	"github.com/TobiEiss/aranGoDriver"
	"log"
	"strconv"
)

var session aranGoDriver.Session

//https://github.com/TobiEiss/aranGoDriver
func handleDatabase() {

	// Initialize a arango-Session with the address to your arangoDB.
	//
	// If you write a test use:
	// session = aranGoDriver.NewTestSession()
	//
	session = aranGoDriver.NewAranGoDriverSession("http://localhost:8529")

	// Connect to your arango-database:
	session.Connect("root", "1234")

	// Concrats, you are connected!
	// Let's print out all your databases
	list, err := session.ListDBs()
	if err != nil {
		log.Fatal("there was a problem: ", err)
	}
	log.Println(list)

	// Create a new database
	err = session.CreateDB("DataVault")
	// TODO: handle err

	// Create a new collection
	err = session.CreateCollection("DataVault", "Data")
	// TODO: handle err

}

func appendToDB(tpackage tempData) {

	foods := map[string]interface{}{
		"Nr":       strconv.Itoa(tpackage.Nr),
		"Speed":    strconv.Itoa(tpackage.Speed),
		"Setpoint": strconv.Itoa(tpackage.Setpoint),
		"Pressure": strconv.Itoa(tpackage.Pressure),
		"Auto":     strconv.FormatBool(tpackage.Auto),
		"Err":      strconv.FormatBool(tpackage.Err),
	}

	_, err := session.CreateDocument("DataVault", "Data", foods)
	if err != nil {
		log.Println(err)
	}

}
