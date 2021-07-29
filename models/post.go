package models

import "time"

// Post 帖子结构体
type Post struct {
	ID       	int64 `json:"id" db:"post_id"`
	AuthorID 	int64 `json:"author_id" db:"author_id"`
	CommunityID int64  `json:"community_id" db:"community_id" binding:"required"` //  binding:"required" 检测发现空值就会报错  这个字段的值必须要有
	Status      int32  `json:"status" db:"status"`
	Title       string `json:"title" db:"title" binding:"required"`
	Content string `json:"content" db:"content" binding:"required"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*Post  // 嵌入帖子结构体
	*CommunityDetail  // 嵌入社区详情结构体
}
