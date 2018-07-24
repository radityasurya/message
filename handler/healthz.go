package handler

import (
	"net/http"
)

const (
	msgOk   = `{"message": "OK"}`
	msgDown = `{"message": "Failed"}`
)

// Healthz will return whether the service and it's dependecies
// is still healthy.
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// TODO: Ping https://rest.messagebird.com/messages

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msgOk))
}
