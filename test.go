package main

import (
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	url                  = "localhost"
	database, collection = "adeptio", "blocks"
)

type (
	JsonAggregation struct {
		Block         int    `json:"block"`
		Hash          string `json:"hash"`
		Confirmations int    `json:"confirmations"`
		Time          int    `json:"time"`
	}
)

var results []interface{}
var orderItems []JsonAggregation

func main() {
	mongoSession()

	//query := coll.Find(nil).All(&results)
	//query := coll.Find(bson.M{"block": 1}).All(&results)
	//query := coll.Find(bson.D{ {"block", bson.D{{"$lt", 100000}}} }).All(&results)
	//query := coll.Find(bson.D{ {"block", 1} }).All(&results)

	// Aggregate query
	jsonString, err := json.Marshal(results)
		if err != nil {
			panic("FATAL, got Error while getting the results from MongoDB")
		} else if string(jsonString) == "null" {
			panic("FATAL, returned empty data from MongoDB (null)")
		} else {

			err := json.Unmarshal(jsonString, &orderItems)
			if err != nil {
				panic("FATAL")
			}
			assignJsonValues()
			}
		}

	func mongoSession() {
		// Main Connect dial;
		session, err := mgo.Dial(url)
		if err != nil {
			panic("MongoDB not working?")
		}
		// Close socket at the end of the process;
		defer session.Close()

		startColBlocksQuery("block", 1, session)
	}

	func startColBlocksQuery(name string, value int, session *mgo.Session) {
		db := session.DB(database)
		coll := db.C(collection)
		query := coll.Find(bson.D{ {name, value} }).All(&results)
		if query != nil {
			panic("No content in DB")
		}
	}

	func assignJsonValues() {
		for _, orderItem := range orderItems {
		hash := orderItem.Hash
		block := orderItem.Block
		confirmations := orderItem.Confirmations
		time := orderItem.Time
		fmt.Println(hash)
		fmt.Println(block)
		fmt.Println(confirmations)
		fmt.Println(time)
	}
}