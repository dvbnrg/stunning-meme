package logger

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Event struct {
	Time          time.Time
	Method        string
	RequestURI    string
	Name          string
	Execution     time.Duration
	Pid           int
	Hostname      string
	UserAgent     string
	Host          string
	RemoteAddress string
	// UserAgentr    string
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		pid := os.Getpid()

		hostname, _ := os.Hostname()

		log.Printf(
			"%s\t%s\t%s\t%s\t%v\t%s\t%s\t%v\t%v",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start).String(),
			pid,
			hostname,
			r.Header.Get("User-Agent"),
			r.Host,
			// r.RemoteAddr,
			r.UserAgent(),
			// w.Header,
			// r.Response.Body,
		)

		e := Event{
			Time:       start,
			Method:     r.Method,
			RequestURI: r.RequestURI,
			Name:       name,
			Execution:  time.Since(start),
			Pid:        pid,
			Hostname:   hostname,
			Host:       r.Host,
			// RemoteAddress: r.RemoteAddr,
			UserAgent: r.UserAgent(),
		}

		PersistToMongo(e)
	})
}

func PersistToMongo(e Event) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://justdave:supersecret@cluster0-xsmx6.gcp.mongodb.net/test?retryWrites=true&w=majority")
	mgo, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := mgo.Database("Logs").Collection("Logs")
	_, err = collection.InsertOne(context.TODO(), e)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("Inserted a Single Document: ", insertResult.InsertedID)
}
