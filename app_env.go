package saddoug

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sadDoug/store"
)

// AppEnv TBD
type AppEnv struct {
	MessageStore store.Message
	Logger       *log.Logger
}

// NewAppEnv TBD
func NewAppEnv() (*AppEnv, error) {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	messageStore, err := store.NewDBMessage(db)
	if err != nil {
		return nil, err
	}

	return &AppEnv{
		Logger:       log.New(os.Stdout, "logger: ", log.Lshortfile),
		MessageStore: messageStore,
	}, nil
}
