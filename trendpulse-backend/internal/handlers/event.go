package handlers

import (
	"net/http"
)

func SubmitEvent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Event received (stub)"))
}
