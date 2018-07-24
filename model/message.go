package model

import "errors"

// Message struct is the message structure provided by Message Bird
type Message struct {
	Recipients int    `json:"recipients"`
	Originator string `json:"originator"`
	Body       string `json:"body"`
}

// Validate checks the message information is empty or not
func (m *Message) Validate() error {
	if m.Recipients == 0 {
		return errors.New("recipient is empty")
	}
	if m.Originator == "" {
		return errors.New("originator is empty")
	}
	if m.Body == "" {
		return errors.New("message is empty")
	}

	return nil
}
