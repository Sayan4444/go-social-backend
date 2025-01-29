package models

import "time"

type Follower struct {
	UserID     uint
	FollowerID uint
	CreatedAt  time.Time
}
