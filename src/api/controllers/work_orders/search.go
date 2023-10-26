package workorderscontrollers

import (
	"github.com/JuanDArdilaG/work-orders-service/src/api/controllers"
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	"github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/apperrors"
	_ "github.com/JuanDArdilaG/work-orders-service/src/contexts/shared/infrastructure/apperrors"
	_ "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/domain"
	workordersusecases "github.com/JuanDArdilaG/work-orders-service/src/contexts/work_orders/usecases"
	"github.com/gin-gonic/gin"
)

type SearchParams struct {
	since  string
	until  string
	status string
}

// Search godoc
// @Summary    	 Search work orders by date range and status
// @Tags         work_orders
// @Param        since   query      string  false  "Since date"
// @Param        until   query      string  false  "Until date"
// @Param        status  query      string  false  "Status"
// @Success      200
// @Failure      400  {object}  	apperrors.JSONError
// @Failure      404  {object}  	apperrors.JSONError
// @Failure      500  {object}  	apperrors.JSONError
// @Router       /workorders/search [get]
func Search(repos dependencies.Repositories) func(c *gin.Context) {
	return func(c *gin.Context) {
		params, err := params(c)
		if err != nil {
			controllers.CatchError(c, err)
			return
		}
		usecase := workordersusecases.NewSearch(repos.WorkOrders)
		orders, err := usecase.Execute(params.since, params.until, params.status)
		if err != nil {
			controllers.CatchError(c, err)
			return
		}
		c.JSON(200, gin.H{"orders": orders})
	}
}

func params(c *gin.Context) (SearchParams, error) {
	since := c.Query("since")
	until := c.Query("until")
	status := c.Query("status")
	if since == "" && until == "" && status == "" {
		return SearchParams{}, apperrors.NewBadRequest("at least one of these parameters must be provided", "since, until, status")
	}
	return SearchParams{
		since:  since,
		until:  until,
		status: status,
	}, nil
}
