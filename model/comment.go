package model

import (
	"time"
)

type Comment struct {
	CommentId string    `json:"commentId"`
	AuthorId  string    `json:"authorId"` // Foreign key to User
	PostId    string    `json:"postId"`   // Foreign key to Post
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
