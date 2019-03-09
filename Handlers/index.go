package handlers

import (
	"fmt"
	"net/http"
)

// IndexHandler - returns hi
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world")
}
