package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Person - is the memory representation of one user person
type Person struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Age   int    `json:"age,omitempty"`
	City  string `json:"city,omitempty"`
}

var people []Person

// GetPersons - Returns all the profile in the Persons Collection
func insertPerson(db *mgo.Session, person *Person) error {
	c := db.DB(*mgo).C("entries")
	count, err := c.Find(bson.M{"id": person.Id}).Limit(1).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("id %s already exists", person.Id)
	}
	return c.Insert(person)
	person := &Person{
		Id:    id,
		Name:  name,
		Phone: phone,
		Age:   age,
		City:  city,
	}

	err = insertPerson(db, person)
	if err != nil {
		log.Println("Could not save the entry to MongoDB:", err)
	}
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people/{id}", insertPerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":27017", router))

}
