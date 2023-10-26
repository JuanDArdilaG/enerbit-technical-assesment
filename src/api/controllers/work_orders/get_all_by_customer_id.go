package workorderscontrollers

import (
	"github.com/JuanDArdilaG/work-orders-service/src/api/controllers"
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	customers "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/domain"
	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/apperrors"
	workordersusecases "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/usecases"
	"github.com/gin-gonic/gin"
)

// Get All By CustomerID
// @Summary    	 Get all work orders by customerID
// @Tags         work_orders
// @Param        id   path      string  true  "Customer ID"
// @Success      200
// @Failure      400  {object}  apperrors.JSONError
// @Failure      404  {object}  apperrors.JSONError
// @Failure      500  {object}  apperrors.JSONError
// @Router       /workorders/customer/{id} [get]
func GetAllByCustomerID(repos dependencies.Repositories) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := customers.NewID(c.Param("id"))
		if err != nil {
			controllers.CatchError(c, apperrors.NewBadRequest("customer.id", c.Param("id")))
			return
		}
		usecase := workordersusecases.NewGetAllByCustomerID(repos.WorkOrders)
		orders, err := usecase.Execute(id)
		if err != nil {
			controllers.CatchError(c, err)
			return
		}
		c.JSON(200, gin.H{"orders": orders})
	}
}
