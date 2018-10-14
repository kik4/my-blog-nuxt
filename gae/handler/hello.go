package handler

import (
	"fmt"
	"net/http"
)

// Hello is handler returns "Hello."
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello.")
}
