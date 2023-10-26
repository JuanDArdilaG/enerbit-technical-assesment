package workorderscontrollers

import (
	"github.com/JuanDArdilaG/work-orders-service/src/api/controllers"
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	customersusecases "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/application"
	workordersusecases "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/usecases"
	"github.com/gin-gonic/gin"
)

// Create Work Order godoc
// @Summary      Create a work order
// @Tags         work_orders
// @Param		 message body CreateRequest true "Work Order Info"
// @Success      200
// @Failure      400  {object}  apperrors.JSONError
// @Failure      404  {object}  apperrors.JSONError
// @Failure      500  {object}  apperrors.JSONError
// @Router       /workorders [post]
func Create(repos dependencies.Repositories) func(c *gin.Context) {
	return func(c *gin.Context) {
		orderBody, err := ParseBody(c)
		if err != nil {
			controllers.CatchError(c, err)
			return
		}
		usecase := workordersusecases.NewCreate(
			repos.WorkOrders,
			repos.Customers,
			customersusecases.NewDeactivate(repos),
		)
		err = usecase.Execute(
			orderBody.CustomerID,
			orderBody.Title,
			orderBody.PlannedDateBegin,
			orderBody.PlannedDateEnd,
		)
		if err != nil {
			controllers.CatchError(c, err)
			return
		}
	}
}
