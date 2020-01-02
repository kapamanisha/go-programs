package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type Person struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Age   int    `json:"age,omitempty"`
	City  string `json:"city,omitempty"`
}

var people []Person

// func GetPerson(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	for _, item := range people {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)

// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Person{})
// }

// func GetPeople(w http.ResponseWriter, req *http.Request) {
// 	json.NewEncoder(w).Encode(people)

// }
func CreatePerson(w http.ResponseWriter, req *http.Request) {
	var database string = "db"

	Session, err := mgo.Dial("mongodb://localhost:27017/" + database)
	if err != nil {
		fmt.Println(err)
	}
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
	Session.person.Insert(people)

}

// func DeletePerson(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	for index, item := range people {
// 		if item.ID == params["id"] {
// 			people = append(people[:index], people[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(people)

// }

var database string = "db"

func main() {

	Session, err := mgo.Dial("mongodb://localhost:27017/" + database)
	if err != nil {
		fmt.Println(err)
	}

	// Cleanup
	defer Session.Close()

	router := mux.NewRouter()
	// router.HandleFunc("/people", GetPeople).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":27017", router))
}
