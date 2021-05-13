package models

import "time"

// 内存对齐概念
type Post struct {
	Title       string    `json:"title" db:"title" bingding:"required"`
	Content     string    `json:"content" db:"content" bingding:"required"`
	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" bingding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*Post
	*CommunityDetail `json:"community"`
}
