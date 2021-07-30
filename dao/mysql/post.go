package mysql

import (
	"bluebellAPI/models"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post)(err error){
	sqlStr := `INSERT INTO post(post_id, title, content, author_id, community_id) VALUES(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostList 根据分页获取帖子详情
func GetPostList(page, size int64) (posts []*models.Post, err error){
	sqlStr := `SELECT post_id, title, content, author_id, community_id, create_time FROM post limit ?,?`
	//fmt.Println(page, size)
	//fmt.Println(sqlStr)
	posts = make([]*models.Post, 0, 2)  // 不要写成 make([]*models.Post, 2)
	err = db.Select(&posts, sqlStr, (page - 1) * size, size)
	//fmt.Println(err)
	//fmt.Println(posts)
	return
}

// GetPostById 根据id查找帖子
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `SELECT 
		post_id, title, content, author_id, community_id, create_time
		FROM 
		post
		WHERE 
		post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}