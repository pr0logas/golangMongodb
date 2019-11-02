package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	//log "github.com/sirupsen/logrus"
	"encoding/json"
)

const (
	url        = "localhost"
	database   = "adeptio"
)

func main() {
	// Main Connect dial;
	session, err := mgo.Dial(url)
	if err != nil {
		panic("Database not working?")
	}

	// Close socket at the end of the process;
	defer session.Close()
	fmt.Printf("Succesfully connected to mongoDB. Server endpoint: %v", url)

	// Get DatabaseNames;
	/*dbName, err := session.DatabaseNames()
	if err != nil {
		log.Warn(err)
	}

	// Loop database names and print out'
	for _, v := range dbName {
		fmt.Printf("\n[%3v]", v)
	}*/

	// Connect to specified databaseName from const;
	db := session.DB(database)

	// Get Collections;
	//collection, err := db.CollectionNames()

	var results []interface{}

	coll := db.C("blocks")

	//query := coll.Find(nil).All(&results)
	//query := coll.Find(bson.M{"block": 1}).All(&results)
	//query := coll.Find(bson.D{ {"block", bson.D{{"$lt", 100000}}} }).All(&results)
	query := coll.Find(bson.D{ {"block", 1} }).All(&results)

    if query != nil {
        // TODO: Do something about the error
    } else {
        //fmt.Println("Results All: ", results) 
        jsonString, err := json.Marshal(results)
		fmt.Println(err)
		fmt.Println(string(jsonString))
    }

	// Loop and printOur collections;
	/*fmt.Printf("\nCollections:\n")
	for _, v:= range collection {
		fmt.Printf("\n[%3v]", v)
	}
	*/



}
