package model

import (
	"strings"
	"time"
)

const (
	// MessageTreeSeparator TBD
	MessageTreeSeparator = ">->"
)

// Message TBD
type Message struct {
	ID int64 `db:"id"`

	Resource string `db:"resource"`
	Topic    string `db:"topic"`

	ExternalID string    `db:"external_id"`
	AnsweredTo string    `db:"answered_to"`
	PostedTime time.Time `db:"posted_time"`

	Author string `db:"author"`

	TreePath string `db:"tree_path"`
}

// JoinMessagePaths TBD
func JoinMessagePaths(paths ...string) string {
	return strings.Join(paths, MessageTreeSeparator)
}
