package handlers

import (
	"fmt"
	"net/http"
)

func PutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Inside the /put handler!")
}