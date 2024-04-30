package model

import (
	"time"
)

type Message struct {
	MessageId      string    `json:"messageId"`
	SenderUserId   string    `json:"senderUserId"`   // Foreign key to User
	ReceiverUserId string    `json:"receiverUserId"` // Foreign key to User
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"createdAt"`
}
