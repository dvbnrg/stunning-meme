package handlers

import (
	"fmt"
	"net/http"
	// "github.com/dvbnrg/zaboCodingChallenge/model"
	// "github.com/dvbnrg/zaboCodingChallenge/db"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "It works!")
}
