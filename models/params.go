package models

//定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username   string `json:"username" binging:"required"`
	Password   string `json:"password" binging:"required"`
	RePassword string `json:"re_password" binging:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binging:"required"`
	Password string `json:"password" binging:"required"`
}

type ParamVoteData struct {
	//UserID 从请求中获取当前的用户
	PostID    string `json:"post_id,string" binging:"required"`
	Direction int8   `json:"direction,string" binging:"oneof=1 0 -1"` //赞成票+1 反对票-1 取消投票0
}

type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"` //可以为空
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
	Order       string `json:"order" form:"order"`
}

//type ParamCommunityPostList struct {
//	*ParamPostList
//	CommunityID int64 `json:"community_id" form:"community_id"`
//}
