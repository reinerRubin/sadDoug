package saddoug

import (
	"log"

	"github.com/sadDoug/model"
	"github.com/sadDoug/store"
)

// MessageSaver TBD
type MessageSaver struct {
	inMessages   messageSource
	messageStore store.Message

	logger *log.Logger
}

// NewMessageSaver TBD
func NewMessageSaver(appEnv *AppEnv, ms messageSource) (*MessageSaver, error) {
	return &MessageSaver{
		logger:       appEnv.Logger,
		inMessages:   ms,
		messageStore: appEnv.MessageStore,
	}, nil
}

// Run TBD
func (ms *MessageSaver) Run() error {
	go func() {
		inMessages := ms.inMessages
		for {
			select {
			case message, opened := <-inMessages:
				if !opened {
					inMessages = nil
					break
				}
				if err := ms.saveMessage(message); err != nil {
					ms.logger.Printf("save msg(%d) err: %s",
						message.ID, err.Error())
				}
			}
		}
	}()

	return nil
}

func (ms *MessageSaver) saveMessage(message *model.Message) error {
	return ms.messageStore.Save(message)
}
