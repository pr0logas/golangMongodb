package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"encoding/json"
)

const (
	url        	   = "localhost"
	databaseName   = "adeptio"
	collectionName = "blocks"
)

func main() {
	// Main Connect dial;
	session, err := mgo.Dial(url)
	if err != nil {
		panic("Database not working?")
	}

	// Close socket at the end of the process;
	defer session.Close()
	//fmt.Printf("Succesfully connected to mongoDB. Server endpoint: %v", url)

	// Connect to specified databaseName from const;
	db := session.DB(databaseName)
		if err != nil {
		fmt.Printf("Failed to connect to specific database %v:Database not working?", databaseName)
		panic("FATAL")
	}

	var results []interface{}

	coll := db.C(collectionName)

	//query := coll.Find(nil).All(&results)
	//query := coll.Find(bson.M{"block": 1}).All(&results)
	//query := coll.Find(bson.D{ {"block", bson.D{{"$lt", 100000}}} }).All(&results)
	query := coll.Find(bson.D{ {"block", 1} }).All(&results)

	// Check if query success;
    if query != nil {
        fmt.Printf("Failed to apply query to specific database %v:Database not working?", databaseName)
		panic("FATAL")
    } else {
        jsonString, err := json.Marshal(results)
        if err != nil {
	        fmt.Printf("Failed to return and convert query to JSON from specific database %v:Database not working?", databaseName)
			panic("FATAL")
        }
		fmt.Println(string(jsonString))
    }
}
