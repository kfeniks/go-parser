package main

import (
	"github.com/go-bongo/bongo"
	"log"
)

type Person struct {
	bongo.DocumentBase `bson:",inline"`
	FirstName string
	LastName string
	Gender string
	HomeAddress struct {
		Street string
		Suite string
		City string
		State string
		Zip string
	}
}

func main() {

	config := &bongo.Config{
		ConnectionString: "user:pass@ip",
		Database:         "parser_email",
	}

	connection, err := bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}

	savePerson(connection)

}

func savePerson(connection *bongo.Connection)  {
	myPerson := &Person{
		FirstName:"Alex",
		LastName:"Dubinin",
		Gender:"male",
	}

	errSave := connection.Collection("people").Save(myPerson)

	if errSave != nil {
		log.Fatal(errSave)
	}
}

