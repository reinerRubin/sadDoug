package saddoug

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sadDoug/model"
)

const (
	// BNWThreadURLTemplate TBD
	BNWThreadURLTemplate = "https://bnw.im/api/show?message=%s&replies=1"
)

type (
	// BNWThread TBD
	BNWThread struct {
		MsgID   string        `json:"msgid"`
		Replies []*BNWMessage `json:"replies"`
		Message *BNWMessage
	}

	// BNWMessage TBD
	BNWMessage struct {
		ID      string  `json:"id"`
		ReplyTo *string `json:"replyto"`

		Date float32 `json:"date"`
		User string  `json:"user"`
	}
)

func (bm *BNWMessage) toMessage() *model.Message {
	unixTime := int64(bm.Date)
	createTime := time.Unix(unixTime, 0)

	return &model.Message{
		Resource:   "bnw.im",
		ExternalID: bm.ID,
		AsnweredTo: bm.ReplyTo,
		Time:       createTime,
		Author:     bm.User,
	}
}

// NewBNWThread TBD
func NewBNWThread(threadID string) (messageSource, error) {
	resp, err := http.Get(fmt.Sprintf(BNWThreadURLTemplate, threadID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bnwThread := &BNWThread{}
	if err := json.Unmarshal(body, bnwThread); err != nil {
		return nil, err
	}

	return BNWThreadToMessageSource(bnwThread)
}

// BNWThreadToMessageSource TBD
func BNWThreadToMessageSource(thread *BNWThread) (messageSource, error) {
	messageChan := make(chan *model.Message)

	go func() {
		dumpBNWThreadToChan(messageChan, thread)
	}()

	return messageChan, nil
}

func dumpBNWThreadToChan(messageChan chan *model.Message, thread *BNWThread) {
	defer close(messageChan)

	messages := make([]*BNWMessage, 0, 1+len(thread.Replies))
	messages = append(messages, thread.Message)

	for _, replyBNWMessage := range thread.Replies {
		messages = append(messages, replyBNWMessage)
	}

	// XXX TODO recursion to fix messagesTreePath
	messagesTreePath := ""
	for _, BNWMessage := range messages {
		message := BNWMessage.toMessage()
		message.Topic = thread.MsgID

		if messagesTreePath != "" {
			messagesTreePath = model.JoinMessagePaths(messagesTreePath, message.ExternalID)
		} else {
			messagesTreePath = message.ExternalID
		}

		message.TreePath = messagesTreePath

		messageChan <- message
	}
}
