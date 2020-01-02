package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Age   int    `json:"age,omitempty"`
	City  string `json:"city,omitempty"`
}

func bootstrap(s *mgo.Session) *mgo.Collection {
	s.DB(database).DropDatabase()
	c := s.DB(database).C("people")
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		Background: true,
	}

	c.EnsureIndex(index)

	return c
}
func create(c *mgo.Collection) {

	err := c.Insert(Person{})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n\n")
}
func read(c *mgo.Collection) {

	// Read Once
	result := Person{}
	query := c.Find(bson.M{"id"})
	err := query.One(&result)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	// Read alot
	var results []Person
	query = c.Find(bson.M{}).Sort("-city")
	err = query.All(&results)

	for _, value := range results {
		fmt.Println(value)
	}
}
func update(c *mgo.Collection) {

	filter := bson.M{"id"}
	err := c.Update(filter)
	if err != nil {
		fmt.Println(err)
		return
	}

	read(c)
	fmt.Printf("\n\n")
}

func delete(c *mgo.Collection) {

	filter := bson.M{"name": ""}
	err := c.Remove(filter)

	if err != nil {
		fmt.Println(err)
		return
	}

	read(c)
	fmt.Printf("\n\n")
}

func main() {

	var database string = "go-development"

	c := bootstrap(session)
	create(c)
	read(c)
	update(c)
	delete(c)

	session, err := mgo.Dial("mongodb://localhost:27017/" + database)
	if err != nil {
		fmt.Println(err)
	}

	// Cleanup
	defer session.Close()

	router := mux.NewRouter()
	router.HandleFunc("/people", Getread).Methods("GET")
	router.HandleFunc("/people/{id}", Getupdate).Methods("GET")
	router.HandleFunc("/people/{id}", create).Methods("POST")
	router.HandleFunc("/people/{id}", delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":27017", router))
}
