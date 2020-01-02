package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

var database string = "crud"

type Person struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	Phone string        `json:"phone" bson:"phone"`
	Age   int           `json:"age" bson:"age"`
	City  string        `json:"city" bson:"city"`
}

session, err := mgo.Dial("mongodb://" + os.Getenv("RATING_DB_HOST"))

if err != nil {
	fmt.Println(err)
}

func main() {
	r := gin.Default()
	r.GET("/add", handler.Create())
	r.Run()


	// Cleanup
	 //defer session.Close()
}

func Create(c *gin.Context) {
	u := Person{
		Name:  c.PostForm("name"),
		Phone: c.PostForm("phone"),
		Age,_:   strconv.Atoi(c.PostForm("age")),
		City:  c.PostForm("city"),
	}


}
