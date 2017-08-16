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
func (d *DBMessage) Save(message *model.Message) error {
	rows, err := d.db.NamedQuery(`
INSERT INTO message (
  resource, topic, external_id, answered_to, posted_time, author, tree_path
) VALUES (:resource, :topic, :external_id, :answered_to, :posted_time, :author, :tree_path)
RETURNING id
        `, message)

	if err != nil {
		panic(err)
	}

	var ID int64
	if rows.Next() {
		if err := rows.Scan(&ID); err != nil {
			return err
		}
	}

	message.ID = ID

	return err
}
