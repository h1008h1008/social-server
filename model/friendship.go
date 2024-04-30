package model

import (
	"time"
)

type Friendship struct {
	UserId1   string    `json:"userId1"` // Foreign key to User
	UserId2   string    `json:"userId2"` // Foreign key to User
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
