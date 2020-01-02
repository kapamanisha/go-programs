package main

import (
	"fmt"

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

var database string = "User"

func main() {

	session, err := mgo.Dial("mongodb://localhost:27017/" + database)
	if err != nil {
		fmt.Println(err)
	}

	// Cleanup
	defer session.Close()

	c := bootstrap(session)

	create(c)
}
// 	read(c)
// 	update(c)
// 	delete(c)
// }

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
	var map
	u := Person{
		ID:    bson.NewObjectid(),
		Name:  c.P["name"].string,
		Phone: payload["phone"].string,
		Age:   payload["age"].int,
		City:  payload["city"].string,
	}
	err := c.Insert(u)

	if err != nil {
		fmt.Println(err)
	}

	
}
// func read(c *mgo.Collection) {

// 	// Read Once
// 	result := Person{}
// 	query := c.Find(bson.M{"name": "Eli"})
// 	query = query.Select(bson.M{"phone": 0})
// 	err := query.One(&result)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(result)

// 	// Read alot
// 	var results []Person
// 	query = c.Find(bson.M{}).Sort("-city")
// 	err = query.All(&results)

// 	for _, value := range results {
// 		fmt.Println(value)
// 	}
// }
// func update(c *mgo.Collection) {

// 	filter := bson.M{"name": "Eli"}
// 	change := bson.M{
// 		"$set": bson.M{
// 			"phone": "+86 99 8888 7777",
// 			"city":  "kerla",
// 		},
// 	}

// 	err := c.Update(filter, change)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	read(c)
// 	fmt.Printf("\n\n")
// }

// func delete(c *mgo.Collection) {

// 	filter := bson.M{"name": "Ben"}
// 	err := c.Remove(filter)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	read(c)
// 	fmt.Printf("\n\n")
// }
