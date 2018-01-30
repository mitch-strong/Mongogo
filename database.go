package database

import (
	"fmt"
	"log"

	"github.com/ljdursi/candig_mds/models"
	"gopkg.in/mgo.v2"
)

var databaseName string
var connectionString string

var session *mgo.Session
var err error

//DatabaseInit creates a connection to the database
func DatabaseInit(dbName, connectionstring string) {
	databaseName = dbName
	connectionString = connectionstring + dbName

	session, err = mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
}

func setCollection(collection string) *mgo.Collection {
	return session.DB(databaseName).C(collection)
}

//Insert allows users to add generic objects to a collection in the database
func Insert(collection string, object interface{}) bool {
	c := setCollection(collection)
	err := c.Insert(object)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// //GetAll returns an array of all objects in a collection
// func GetAll(collection string) []interface{} {
// 	c := setCollection(collection)
// 	var list []interface{}
// 	err := c.Find(nil).All(&list)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return list
// }

func GetAllBiosamples(collection string) []*models.Biosample {
	c := setCollection(collection)
	var list []*models.Biosample
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

func GetAll(collection string) []interface{} {
	var result interface{}
	c := setCollection(collection)
	iter := c.Find(nil).Iter()
	count, _ := c.Find(nil).Count()
	list := make([]interface{}, count)
	counter := 0
	for iter.Next(&result) {
		list[counter] = result
		counter++
	}
	if err := iter.Close(); err != nil {
		fmt.Printf("%v", err)
	}
	return list
}

//RemoveAll will empty a collection
func RemoveAll(collection string) bool {
	c := setCollection(collection)
	_, err := c.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// //Test tests connection with simple queries
// func Test() {
// 	var result interface{}
// 	c := setCollection("objects")
// 	err = c.Find(bson.M{"first": "mitchell"}).One(&result)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//fmt.Printf("%v", result)
// 	addPerson()
// 	addPerson()
// 	addPerson()

// 	fmt.Printf("%v", GetAll("people"))
// 	RemoveAll("people")
// }

// func addPerson() {
// 	pers := Person{FirstName: "Mitchell", LastName: "Strong", IsUHN: true, AddedOn: time.Now()}
// 	Insert("people", pers)
// }
