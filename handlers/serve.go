package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dvbnrg/zaboCodingChallenge/db"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	x := db.Serve()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(x); err != nil {
		panic(err)
	}
}
