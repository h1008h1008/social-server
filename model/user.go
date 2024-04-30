package model

import (
	"time"
)

type User struct {
	UserId            string    `json:"userId"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordHash      string    `json:"passwordHash"`
	ProfilePictureUrl string    `json:"profilePictureUrl"`
	Bio               string    `json:"bio"`
	CreatedAt         time.Time `json:"createdAt"`
}
