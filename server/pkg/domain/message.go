// Message type for the message queue
package domain

import (
	"github.com/google/uuid"
)

type Message struct {
	UUID       uuid.UUID
	Author     string
	Content    string
	ParentUUID uuid.UUID
	TimeStamp  int64
}

type MessageList struct {
	Messages []Message
}
