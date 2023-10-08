package handlers

import (
	"fmt"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Inside the /get handler!")
}