package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID    bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string        `bson:"name,omitempty" json:"name,omitempty"`
	Phone string        `bson:"phone,omitempty" json:"phone,omitempty"`
	Age   string        `bson:"age,omitempty" json:"age,omitempty"`
	City  string        `bson:"city,omitempty" json:"_city,omitempty"`
}

var database string = "db"

func main() {

	session, err := mgo.Dial("mongodb://localhost:27017/" + database)
	if err != nil {
		fmt.Println(err)
	}

	// Cleanup
	defer session.Close()
	r := gin.Default()
	c := bootstrap(session)
	r.POST("/add", func(data *gin.Context) {
		create(c, data)
	})
	r.GET("/read", func(data *gin.Context) {
		//data.JSON(200, gin.H{
		//"data": result,
		//})
		read(c, data)

	})
	r.DELETE("/delete", func(data *gin.Context) {
		delete(c, data)
	})
	r.PUT("/update", func(data *gin.Context) {
		update(c, data)
	})
	r.Run()
}
func bootstrap(s *mgo.Session) *mgo.Collection {

	s.DB(database) //.DropDatabase()
	c := s.DB(database).C("people")
	return c
}

func create(c *mgo.Collection, data *gin.Context) {
	// payload := make(map[string]interface{},0)
	fmt.Println("hello manisha")

	person := Person{
		ID:    bson.NewObjectId(),
		Name:  data.PostForm("name"),
		Phone: data.PostForm("phone"),
		Age:   data.PostForm("age"),
		City:  data.PostForm("city"),
	}
	err := c.Insert(person)
	if err != nil {
		fmt.Println(err)

	}

	data.JSON(200, gin.H{
		"status":  "posted",
		"message": person,
	})

}
func read(c *mgo.Collection, data *gin.Context) []Person {
	fmt.Println("hello manisha")
	// Read alot
	var result []Person
	query := c.Find(bson.M{})
	//query = query.Select(bson.M{})
	err := query.All(&result)
	if err != nil {
		fmt.Println(err)
	}
	// var results []Person
	// query := c.Find(bson.M{}).Sort("-name")
	// err := query.All(&results)

	// for _, value := range results {
	// 	fmt.Println(value)
	//  }

	//["people"].Find(nil).all(result)
	fmt.Println("hellllllllllllllllllllooooooooooooooo", result)
	//fmt.Printf("\n\n")
	// return result

	data.JSON(200, gin.H{
		"status":  "posted",
		"message": result,
	})
	return result
}
func delete(c *mgo.Collection, data *gin.Context) {
	fmt.Println("deleting")

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(data.PostForm("id"))})
	fmt.Println(err)

	if err != nil {
		fmt.Println(err)
		return
	}
	data.JSON(200, gin.H{
		"status":  "posted",
		"message": "user deleted successfully",
	})

	//read(c)
	fmt.Printf("\n\n")
}
func update(c *mgo.Collection, data *gin.Context) {
	fmt.Println(update)

	filter := bson.M{}
	change := bson.M{
		"$set": bson.M{
			//Name:  data.PostForm("name"),
			//Phone: data.PostForm("phone"),
		},
	}

	err := c.Update(filter, change)
	if err != nil {
		fmt.Println(err)
		return
	}
	//data.JSON(200, gin.H{
	//"status":  "posted",
	//"message": change,
	//})

	//read(c)
	fmt.Printf("\n\n")
}
