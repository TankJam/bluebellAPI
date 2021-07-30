package logic

import (
	"bluebellAPI/dao/mysql"
	"bluebellAPI/models"
	"bluebellAPI/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	// 1.生成post id
	p.ID = snowflake.GenID()
	// 2.保存到数据库，并返回
	return mysql.CreatePost(p)
}


// GetPostList 获取数据
func GetPostList(page, size int64) (data [] *models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err),
			)
			continue
		}

		// 根据社区id查询详情信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err),
			)
		}
		fmt.Println(community)

		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}

		data = append(data, postDetail)
	}
	return
}

// GetPostById 根据id查看帖子详情数据
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	// 1.查询并组合我们接口想用得数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed",
			zap.Int64("pid", pid),
			zap.Error(err),
		)
		return
	}

	// 根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err),
		)
		return
	}

	// 根据社区id查找社区信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err),
		)
		return
	}

	// 接口数据拼接
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}