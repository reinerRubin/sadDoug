package saddoug

import (
	"log"

	"github.com/sadDoug/model"
)

// MessageSaver TBD
type MessageSaver struct {
	logger     *log.Logger
	inMessages messageSource
}

// NewMessageSaver TBD
func NewMessageSaver(appEnv *AppEnv, ms messageSource) (*MessageSaver, error) {
	return &MessageSaver{
		logger:     appEnv.Logger,
		inMessages: ms,
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
	ms.logger.Printf("save msg(%d)", message.ID)
	return nil
}
