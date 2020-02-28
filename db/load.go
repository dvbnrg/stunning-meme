package db

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dvbnrg/zaboCodingChallenge/model"
)

func LoadDataToMgo() {
	resp, err := http.Get("https://api.code.gov/repos?api_key=IiLkBh9KDgZEhlJ6JghDKjKqWKklAZemuMBfxNow")
	if err != nil {
		log.Println("Data Retreival Error: ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	repos := model.Repos{}
	err = json.Unmarshal(body, &repos)
	db := mgoConn()
	collection := db.Database("CodeGovRepos").Collection("Repos")
	// log.Println(string(body))
	insertResult, err := collection.InsertOne(context.TODO(), repos)
	if err != nil {
		log.Println("Data Insertion Error: ", err)
	}
	log.Println("Inserted a Single Document: ", insertResult.InsertedID)
}
