package database

import (
	"go-postbar/logger"
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

	// Optional. Switch the session to a monotonic behavior.
	// session.SetMode(mgo.Monotonic, true)

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
	err = collection.Insert(&model)

	if err != nil {
		log.Fatal(err)
	}

}

// GetAll ...
func GetAll(collectionName string) []interface{} {

	collection := mongo.C(collectionName)

	var results []interface{}

	err = collection.Find(bson.M{}).All(&results)

	if err != nil {
		log.Debug("Mongo retornou um erro: %s", err)
	}

	return results

}

// GetOne ...
func GetOne(collectionName string, id string) interface{} {
	collection := mongo.C(collectionName)

	var result interface{}

	if !bson.IsObjectIdHex(id) {
		log.Error("%s is not a ObjectId value", id)
		return nil
	}

	oid := bson.ObjectIdHex(id)

	err = collection.FindId(oid).One(&result)

	if err != nil {
		log.Debug("Mongo retornou um erro: %s", err)
		return nil
	}

	return result
}

// DeleteOne ...
func DeleteOne(collectionName string, id string) bool {
	collection := mongo.C(collectionName)

	if !bson.IsObjectIdHex(id) {
		log.Error("%s is not a ObjectId value", id)
		return false
	}

	oid := bson.ObjectIdHex(id)

	err = collection.RemoveId(oid)

	if err != nil {
		log.Debug("Mongo retornou um erro: %s", err)
		return false
	}

	return true
}
