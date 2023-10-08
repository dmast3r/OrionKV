package handlers

import (
	"fmt"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Inside the /delete handler!")
}