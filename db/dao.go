package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// dbConn creates an instance of a sql connection then returns the connection details as a struct
// func dbConn() (db *sql.DB) {
// 	dbDriver := "mysql"
// 	dbUser := "justdave"
// 	dbPass := "supersecret"
// 	dbName := "repositories"
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbName)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

func mgoConn() (mgo *mongo.Client) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://justdave:supersecret@cluster0-xsmx6.gcp.mongodb.net/test?retryWrites=true&w=majority")
	mgo, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return
}
