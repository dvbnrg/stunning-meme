package db

import (
	"context"
	"log"

	"github.com/dvbnrg/zaboCodingChallenge/model"
	"go.mongodb.org/mongo-driver/bson"
)

func Serve() (o model.Repos) {
	db := mgoConn()
	collection := db.Database("CodeGovRepos").Collection("Repos")
	filter := bson.D{{}}
	err := collection.FindOne(context.TODO(), filter).Decode(&o)
	if err != nil {
		log.Println("Data Retreival Error: ", err)
	}
	log.Println("Retreived a Single Document")
	return
}
