package pubsubinfra

import "github.com/go-redis/redis"

type RedisStream struct {
	channel string
	Client  *redis.Client
}

func NewRedisStream(channel string, client *redis.Client) *RedisStream {
	return &RedisStream{channel, client}
}

func (rs *RedisStream) Publish(values map[string]interface{}) error {
	return rs.Client.XAdd(&redis.XAddArgs{
		Stream: rs.channel,
		Values: values,
	}).Err()
}
