package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string `bson:"name,omitempty"`
	Phone string `bson:"phone,omitempty"`
	Age   string `bson:"age,omitempty"`
	City  string `bson:"city,omitempty"`
}

var database string = "development"

func main() {

	session, err := mgo.Dial("mongodb://localhost:27017/" + database)
	if err != nil {
		fmt.Println(err)
	}

	// Cleanup
	defer session.Close()

	c := bootstrap(session)

	create(c)
	read(c)
	update(c)
	delete(c)
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
func createAdd(c *gin.Context){
	
	
   //create an object to insert data
	 u := Person{
		 ID:  bson.NewObjectid(),
		 Name:c.PostForm("name"),
		 Phone: c.PostForm("phone"),
		 Age:  c.PostForm("age"),
		 City: c.PostForm("city"),	 
	 }
	 err :=c.Insert(u)
	 

	
	
}
func read(c *gin.Context) {
	var payload map[string]interface{}
	c.BindJSON(&payload)
	// Read Once
	result := Person{}
	query := c.Find(bson.M{"name": []string})
	query = query.Select(bson.M{"phone": []string})
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
func update(c *gin.Context) {
	defer RecoverPanic()
	var payload payload
	c.BindJSON(&payload)

	filter := bson.M{"name": []string}
	change := bson.M{
		"$set": bson.M{
			"phone": []string,
			"city":  []string,
		},
	}

	err := c.Update(filter, change)
	if err != nil {
		fmt.Println(err)
		return
	}

	read(c)
	fmt.Printf("\n\n")
}

func delete(c *gin.Context) {
	defer RecoverPanic()
	var payload payload
	c.BindJSON(&payload)

	filter := bson.M{"name": []string}
	err := c.Remove(filter)

	if err != nil {
		fmt.Println(err)
		return
	}

	read(c)
	fmt.Printf("\n\n")
}
// func (){
// 	// r := gin.Default()
	
// }

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}