package models

import (
	"time"
)

type Message struct {
	ContactName    string    `validation:"required,max=100"`
	ContactEmail   string    `validation:"required,email"`
	MessageSubject string    `validation:"required,max=200"`
	Message        string    `validation:"required,max=500"`
	Timestamp      time.Time `validation:"-"`
}

func (m *Message) FormatTimestamp() string {
	return m.Timestamp.Format("2006-01-02 15:04:05")
}
