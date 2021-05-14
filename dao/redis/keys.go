package redis

//redis key

//redis key尽量使用命名空间的方式，方便查询和拆分

const (
	PreFix            = "bluebell"
	KeyPostTimeZSet   = "post:time"   //zset;贴子及发帖时间
	KeyPostScoreZSet  = "post:score"  //zset;贴子及投票的分数
	KeyPostVoteZSetPF = "post:voted:" //zset；记录用户及投票的类型;参数是post id

	KeyCommunitySetPF = "community:" //set;保存每个分区下帖子的id
)

func getRedisKey(key string) string {
	return PreFix + key
}
