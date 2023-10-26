package customerscontrollers

import (
	"github.com/JuanDArdilaG/work-orders-service/src/api/controllers"
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	customersusecases "github.com/JuanDArdilaG/work-orders-service/src/contexts/customers/application"
	"github.com/gin-gonic/gin"
)

// Get Active Customers
// @Summary    	 Get all active customers
// @Tags         customers
// @Success      200
// @Failure      400  {object}  apperrors.JSONError
// @Failure      404  {object}  apperrors.JSONError
// @Failure      500  {object}  apperrors.JSONError
// @Router       /customers/active [get]
func GetActive(repos dependencies.Repositories) func(c *gin.Context) {
	return func(c *gin.Context) {
		usecase := customersusecases.NewGetActive(repos.Customers)
		customers, err := usecase.Execute()
		if err != nil {
			controllers.CatchError(c, err)
			return
		}
		c.JSON(200, gin.H{"customers": customers})
	}
}
