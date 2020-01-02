package main

import (
	"fmt"
)
type Person struct{
	ID unit `json:"id"`
	Name string `json:"name"`
	Age int    `json:"age"`
	City string `json:"city"`
	Phone no string `json:"phone no"`
}
func main (){
	db, _ := gorm.Open("mongodb","./gorm.db")
	defer db.Close()
	db.AutoMigrate(&Person{})
 p1 := Person{Name: "Kittu",Age: 12,City: "gadwal"}
 p2 := Person{Name: "Kunni",Age: 12,City: "nalgonda"}
 db.Create(&p1)
 var p3 Person // identify a Person type for us to store the results in
 db.City(&p3) // Find the first record in the Database and store it in p3
 fmt.Println(p1.Name)
 fmt.Println(p2.Age)
 fmt.Println(p3.City) // print out our record from the database
}
