package workorderscontrollers

import (
	"github.com/JuanDArdilaG/work-orders-service/src/api/controllers"
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	customersusecases "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/application"
	workorders "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
	workordersusecases "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/usecases"
	"github.com/gin-gonic/gin"
)

// Finish Work Order godoc
// @Summary      Finish a work order
// @Tags         work_orders
// @Param        id   path      string  true  "Order ID"
// @Success      200
// @Failure      400  {object}  apperrors.JSONError
// @Failure      404  {object}  apperrors.JSONError
// @Failure      500  {object}  apperrors.JSONError
// @Router       /workorders/{id}/finish [post]
func Finish(repos dependencies.Repositories, pubs dependencies.Publishers) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := workorders.NewID(c.Param("id"))
		if err != nil {
			controllers.CatchError(c, err)
			return
		}
		usecase := workordersusecases.NewFinish(
			repos.WorkOrders,
			pubs.WorkOrders,
			customersusecases.NewActivate(repos),
		)
		err = usecase.Execute(id)
		if err != nil {
			controllers.CatchError(c, err)
			return
		}
	}
}
