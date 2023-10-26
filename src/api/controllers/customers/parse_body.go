package customerscontrollers

import (
	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/apperrors"
	"github.com/gin-gonic/gin"
)

func ParseBody(c *gin.Context) (customers.Customer, error) {
	customer := customers.Customer{}
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		return customers.Customer{}, apperrors.NewBadRequest("workorder", err.Error())
	}
	err = customer.Validate()
	if err != nil {
		return customers.Customer{}, err
	}
	return customer, nil
}
