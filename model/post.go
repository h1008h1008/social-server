package model

import (
	"time"
)

type Post struct {
	PostId    string    `json:"postId"`
	AuthorId  string    `json:"authorId"` // Foreign key to User
	Content   string    `json:"content"`
	ImageUrls []string  `json:"imageUrls"`
	CreatedAt time.Time `json:"createdAt"`
	Comments  []Comment `json:"comments"` // Embedded Comments
	Likes     []string  `json:"likes"`    // UserIds who liked the post
}
