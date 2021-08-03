package redis

import (
	"bluebellAPI/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

// getIDsForKey 根据起始位置返回结果
func getIDsForKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	// 3. ZREVRANGE 按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

// GetPostIDsInOrder 根据order去获取用户id
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 1.根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	// 2.确定查询的索引起始点
	return getIDsForKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	// 使用pipeline一次发送多条命令，减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIDsInOrder 根据ids获取社区忒子详情
func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 使用 zinterstore 把分区的帖子set与帖子分数的 zset 生成一个新的zset
	// 针对新的zset 按之前的逻辑取数据
	// 社区的key
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))

	// 利用缓存key减少zinterstore的执行次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))

	if client.Exists(key).Val() < 1 {
		// 缓存若不存在，需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",  // 计算最大值
		}, cKey, orderKey)
		pipeline.Expire(key, 60 * time.Second)  // 设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	// 存在的话就直接根据key查询ids
	return getIDsForKey(key, p.Page, p.Size)
}
