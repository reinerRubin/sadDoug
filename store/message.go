package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sadDoug/model"
)

type (
	// Message TBD
	Message interface {
		GetMessageByExternalID(externalID string) (*model.Message, error)
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

// GetMessageByExternalID TBD
func (d *DBMessage) GetMessageByExternalID(externalID string) (*model.Message, error) {
	messages := []model.Message{}
	err := d.db.Select(&messages, "SELECT * FROM message WHERE external_id = $1", externalID)
	if err != nil {
		return nil, err
	}

	if len(messages) > 1 {
		err := fmt.Errorf("multiply messages with uniq id: %s", externalID)
		return nil, err
	}

	var message *model.Message
	if len(messages) == 1 {
		m := messages[0]
		message = &m
	}

	return message, nil
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
	defer func() {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}()

	var ID int64
	if rows.Next() {
		if err := rows.Scan(&ID); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	message.ID = ID

	return err
}
