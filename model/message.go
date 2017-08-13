package model

import (
	"strings"
	"time"
)

const (
	// MessageTreeSeparator TBD
	MessageTreeSeparator = "^^"
)

// Message TBD
type Message struct {
	ID int64

	Resource string
	Topic    string

	ExternalID string
	AsnweredTo *string
	Time       time.Time

	Author string

	TreePath string
}

// JoinMessagePaths TBD
func JoinMessagePaths(a, b string) string {
	return strings.Join([]string{a, b}, MessageTreeSeparator)
}
