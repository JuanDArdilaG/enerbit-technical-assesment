package routes

import (
	workorderscontrollers "github.com/JuanDArdilaG/work-orders-service/src/api/controllers/work_orders"
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	"github.com/gin-gonic/gin"
)

func WorkOrders(repos dependencies.Repositories, pubs dependencies.Publishers, router *gin.RouterGroup) {
	// GET Method
	router.GET("/workorders/customer/:id", workorderscontrollers.GetAllByCustomerID(repos))
	router.GET("/workorders/search", workorderscontrollers.Search(repos))
	// POST Method
	router.POST("/workorders", workorderscontrollers.Create(repos))
	router.POST("/workorders/:id/finish", workorderscontrollers.Finish(repos, pubs))
}
