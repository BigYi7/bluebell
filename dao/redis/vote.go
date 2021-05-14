package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 //每票的分数
)

var (
	ErrVoteTimeExpire = errors.New("超出投票时间")
	ErrVoteRepested   = errors.New("不允许重复投票")
)

func VoteForPost(userID, postID string, value float64) (err error) {
	//1.判断投票限制
	//去redis取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	pipeline := client.TxPipeline()
	//2.更新帖子的分数
	//先查当前用户给当前帖子的之前的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVoteZSetPF+postID), userID).Val()
	//如果这一次和之前保存的一直，就不允许投票
	if value == ov {
		return ErrVoteRepested
	}
	var op float64 //现在的值大于以前的值，分数就是+
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算差值的绝对值
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	//3.记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVoteZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVoteZSetPF+postID), redis.Z{
			Score:  value, //赞成票还是反对票
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	return err

}

func CreatePost(postID, communityID int64) (err error) {

	pipeline := client.TxPipeline()

	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//把帖子ID加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err = pipeline.Exec()
	return err
}
