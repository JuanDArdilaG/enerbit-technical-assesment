package routes

import (
	customerscontrollers "github.com/JuanDArdilaG/work-orders-service/src/api/controllers/customers"
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	"github.com/gin-gonic/gin"
)

func Customers(repos dependencies.Repositories, router *gin.RouterGroup) {
	// GET Method
	router.GET("/customers/active", customerscontrollers.GetActive(repos))
	// POST Method
	router.POST("/customers", customerscontrollers.Create(repos))
}
