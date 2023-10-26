package workordersinfra

import (
	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/pubsubinfra"
	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
)

type RedisPublisher struct {
	stream pubsubinfra.RedisStream
}

func NewRedisPublisher(stream pubsubinfra.RedisStream) workorders.Publisher {
	return &RedisPublisher{stream}
}

func (rp *RedisPublisher) Publish(orderID workorders.ID, status workorders.Status) error {
	return rp.stream.Publish(map[string]interface{}{
		"order_id": orderID,
		"status":   status,
	})
}
