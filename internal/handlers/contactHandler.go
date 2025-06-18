package handlers

import (
	"fmt"
	"project/internal/logging"
	"project/internal/models"
	"project/internal/validation"
	"net/http"
	"time"
)

func ContactHandler(w http.ResponseWriter, req *http.Request) {
	log := logging.Log
	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Error Parsing Form", http.StatusInternalServerError)
	}

	msg := models.Message{
		ContactName:    req.FormValue("name"),
		ContactEmail:   req.FormValue("email"),
		MessageSubject: req.FormValue("subject"),
		Message:        req.FormValue("message"),
		Timestamp:      time.Now(),
	}

	errors := validation.ValidateMessage(msg)
	if len(errors) > 0 {
		log.Errorf("Message Error: %s\nMessage: %+v", errors, msg)

		http.Error(w, "Form Error", http.StatusBadRequest)
		return
	}

	log.Infof("Message: %+v", msg)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<p>Message Received, Thank you %s</p>", msg.ContactName)
}
