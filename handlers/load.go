package handlers

import (
	"net/http"

	"github.com/dvbnrg/zaboCodingChallenge/db"
)

func Load(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	db.LoadDataToMgo()
	// if err := json.NewEncoder(w).Encode(model.Repos); err != nil {
	// 	panic(err)
	// }
}
