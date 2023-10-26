package main

import (
	"net/http"

	_ "github.com/JuanDArdilaG/work-orders-service/docs"
	"github.com/JuanDArdilaG/work-orders-service/src/api/dependencies"
	"github.com/JuanDArdilaG/work-orders-service/src/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Work Orders Service API - EnerBit
// @version         0.1
// @description     This is a sample server for a Work Orders Service API.

// @contact.name   EnerBit

// @host      localhost:8888
// @BasePath  /v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	repositories, err := dependencies.BuildRepositories()
	if err != nil {
		panic(err)
	}
	publishers, err := dependencies.BuildPublishers()
	if err != nil {
		panic(err)
	}
	v1 := r.Group("/v1")
	{
		routes.WorkOrders(*repositories, *publishers, v1)
		routes.Customers(*repositories, v1)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
