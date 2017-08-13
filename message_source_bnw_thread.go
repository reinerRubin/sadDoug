package saddoug

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
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
		ID      string `json:"id"`
		ReplyTo string `json:"replyto"`

		Date float32 `json:"date"`
		User string  `json:"user"`
	}
)

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
	spew.Dump("thread", bnwThread)

	return nil, nil
}
