package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"

	"go.uber.org/zap"
)

//本项目使用简化版的投票分数
//投一票就加432分 86400/200 -> 需要200张赞成票可以给你的帖子续一天 -> redis实战

/*投票的几种情况
direction =1 时,有两种情况：
   1.之前没有投过票，现在投赞成票 -->更新分数和投票记录 差值的绝对值1
   2.之前投反对票，现在改投赞成票 -->更新分数和投票记录 差值的绝对值2
direction =0 时,有两种情况：
   1.之前投过赞成票，现在要取消 -->更新分数和投票记录 差值的绝对值1
   2.之前投过反对票，现在要取消 -->更新分数和投票记录 差值的绝对值1
direction =-1 时,有两种情况：
   1.之前没有投票，现在投反对票 -->更新分数和投票记录 差值的绝对值1
   2.之前投赞成票，现在投反对票 -->更新分数和投票记录 差值的绝对值2

投票的限制：
每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期不允许再投票
  1.到期之后将redis中保存的数据存储到mysql
  2.到期之后删除KeyPostVotePrefix
*/

//为帖子投票的函数
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postid", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
