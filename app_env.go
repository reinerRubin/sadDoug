package saddoug

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	// init pg
	_ "github.com/lib/pq"
	"github.com/sadDoug/store"
)

const (
	// DBAuthStub TBD
	DBAuthStub = "user=saddoug_app dbname=saddoug sslmode=disable password=saddoug_password"
)

// AppEnv TBD
type AppEnv struct {
	MessageStore store.Message
	Logger       *log.Logger
}

// NewAppEnv TBD
func NewAppEnv() (*AppEnv, error) {
	db, err := sqlx.Connect("postgres", DBAuthStub)
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
