package mysql

import (
	"bluebellAPI/models"
	"database/sql"
	"go.uber.org/zap"
)

// GetCommunityList 返回所有社区数据
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		// 若没有返回数据，则日志记录警告
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}


// GetCommunityDetailByID 根据id获取社区详情数据
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `SELECT community_id, community_name, introduction, create_time FROM community where community_id=?`

	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows{
			err = ErrorInvalidID
		}
	}
	return community, err
}

