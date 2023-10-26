package dependencies

import (
	"os"

	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/pubsubinfra"
	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
	workordersinfra "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/infrastructure"
	"github.com/go-redis/redis"
)

type Publishers struct {
	WorkOrders workorders.Publisher
}

func BuildPublishers() (*Publishers, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})
	redisStream := pubsubinfra.NewRedisStream(os.Getenv("REDIS_WORKORDERS_CHANNEL"), redisClient)
	woPublisher := workordersinfra.NewRedisPublisher(*redisStream)
	return &Publishers{
		WorkOrders: woPublisher,
	}, nil
}
