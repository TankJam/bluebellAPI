package models

import "time"

// Community 社区
type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

// CommunityDetail 社区详情
type CommunityDetail struct {
	ID int64 `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
	// omitempty 若数据不存在则不反悔数据
	Introduction string `json:"introduction,omitempty" db:"introduction"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
