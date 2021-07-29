package logic

import (
	"bluebellAPI/dao/mysql"
	"bluebellAPI/models"
)

// GetCommunityList  查询社区列表数据
func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库，查找所有的community，并返回
	return mysql.GetCommunityList()
}
