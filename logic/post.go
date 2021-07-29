package logic

import (
	"bluebellAPI/dao/mysql"
	"bluebellAPI/models"
	"bluebellAPI/pkg/snowflake"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	// 1.生成post id
	p.ID = snowflake.GenID()
	// 2.保存到数据库，并返回
	return mysql.CreatePost(p)
}
