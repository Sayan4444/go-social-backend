package models

import "time"

type Comment struct {
	ID        uint
	PostID    uint
	UserID    uint
	Content   string
	CreatedAt time.Time
}
