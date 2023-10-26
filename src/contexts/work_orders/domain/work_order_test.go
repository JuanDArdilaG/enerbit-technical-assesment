package workorders_test

import (
	"testing"
	"time"

	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
	"github.com/stretchr/testify/assert"
)

func Test_CreateWorkOrderInvalidDates(t *testing.T) {
	t.Run("should return error when planned date end is more than 2 hours forward", func(t *testing.T) {
		id, _ := customers.NewRandomID()
		_, err := workorders.New(
			customers.Customer{ID: id}, "title", time.Now(), time.Now().Add(time.Hour*2+time.Minute),
		)

		assert.Equal(t, workorders.ErrInvalidPlannedDate, err)
	})

	t.Run("should return error when planned date end is before planned date start", func(t *testing.T) {
		id, _ := customers.NewRandomID()
		_, err := workorders.New(
			customers.Customer{ID: id}, "title", time.Now(), time.Now().Add(-time.Hour),
		)

		assert.Equal(t, workorders.ErrInvalidPlannedDate, err)
	})
}
