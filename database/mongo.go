package database

import (
	"fmt"
	"go-postbar/logger"
	"go-postbar/model"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	log     = logger.Log
	session *mgo.Session
	err     error
	mongo   *mgo.Database
)

// Ping execute a ping and return the operation time
func Ping() time.Duration {

	start := time.Now()

	err = session.Ping()
	if err != nil {
		log.Panic(err)
	}

	stop := time.Now()

	return stop.Sub(start)

}

// Dial Conecta no mongo
func Dial(mongoURL string) {
	log.Debug("Calling mongo ...")

	session, err = mgo.Dial(mongoURL)
	mongo = session.DB("postbar")

	if err != nil {
		log.Panic(err)
	}

	elapsed := Ping()

	log.Debug("Mongo ok, response in %s", elapsed)
}

// Close ..
func Close() {
	session.Close()
}

// Save ...
func Save(collectionName string, model interface{}) {
	collection := mongo.C(collectionName)
	err = collection.Insert(model)

	if err != nil {
		log.Fatal(err)
	}

}

// GetAll ...
func GetAll(collectionName string) {

	collection := mongo.C(collectionName)

	var result model.Post

	err = collection.Find(bson.M{"name": "Ale"}).All(&result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Service started, address: %s", result)

	// result := Person{}
	// err =
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

// // Optional. Switch the session to a monotonic behavior.
// session.SetMode(mgo.Monotonic, true)

// c := session.DB("test").C("people")
// err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
// 	&Person{"Cla", "+55 53 8402 8510"})
// if err != nil {
// 	log.Fatal(err)
// }

// result := Person{}
// err = c.Find(bson.M{"name": "Ale"}).One(&result)
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Println("Phone:", result.Phone)
