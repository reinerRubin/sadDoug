package store

import (
	"github.com/jmoiron/sqlx"
	"github.com/sadDoug/model"
)

type (
	// Message TBD
	Message interface {
		Save(*model.Message) error
	}

	// DBMessage TBD
	DBMessage struct {
		db *sqlx.DB
	}
)

// NewDBMessage TBD
func NewDBMessage(db *sqlx.DB) (Message, error) {
	return &DBMessage{
		db: db,
	}, nil
}

// Save TBD
func (d *DBMessage) Save(*model.Message) error {
	return nil
}
