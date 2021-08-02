package models

/*
	定义请求的参数结构体
*/
const (
	OrderTime = "time"
	OrderScore = "score"
)

// ParamSignUp  注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	// required,eqfield=Password RePassword 与 Password 校验
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票参数校验
type ParamVoteData struct {
	// UserID 请求中可以获取
	PostID string `json:"post_id" binding:"required"`  // 贴子id
	// 赞成票(1)还是反对票(-1)取消投票(0)
	Direction int8 `json:"direction,string" binding:"oneof=1 0 -1"`
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64 `json:"community_id" form:"community_id"` // 可以为空
	Page int64 `json:"page" form:"page"`
	Size int64 `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
