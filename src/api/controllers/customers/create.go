package customerscontrollers

import (
	"github.com/JuanDArdilaG/work-orders-service/src/api/controllers"
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	customersusecases "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/application"
	"github.com/gin-gonic/gin"
)

// Create Customer godoc
// @Summary      Create a customer
// @Tags         customers
// @Param		 message body CreateRequest true "Customer Info"
// @Success      200
// @Failure      400  {object}  apperrors.JSONError
// @Failure      404  {object}  apperrors.JSONError
// @Failure      500  {object}  apperrors.JSONError
// @Router       /customers [post]
func Create(repos dependencies.Repositories) func(c *gin.Context) {
	return func(c *gin.Context) {
		customer, err := ParseBody(c)
		if err != nil {
			controllers.CatchError(c, err)
			return
		}
		usecase := customersusecases.NewCreate(repos.Customers)
		err = usecase.Execute(
			customer.FirstName,
			customer.LastName,
			customer.Address,
		)
		if err != nil {
			controllers.CatchError(c, err)
			return
		}
	}
}
