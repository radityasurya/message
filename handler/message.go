package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/radityasurya/message/model"
	log "github.com/sirupsen/logrus"
)

// Message will process the request for handling the
// sending message using MessageBird API.
func Message(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		log.Error(err)
		return
	}
	defer r.Body.Close()

	if string(b) == "" {
		w.WriteHeader(http.StatusNoContent)
		log.Error("No body message")
		return
	}

	msg := model.Message{}
	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Error(err)
		return
	}

	err = msg.Validate()
	if err != nil {
		log.Error(err)
		return
	}

	// Send Message

	log.Info(msg)

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
