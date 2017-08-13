package model

import "time"

// Message TBD
type Message struct {
	ID int64

	Resource string
	Topic    string

	ExternalID string
	AsnweredTo string
	Time       time.Time

	Author string

	TreePath string
}
