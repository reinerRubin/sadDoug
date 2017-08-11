package model

// Message TBD
type Message struct {
	ID int64

	Resource string
	Topic    string

	ExternalID string
	AsnweredTo string

	Author string

	TreePath string
}
