package main

import (
	"assign2/controller"
	"assign2/database"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "assign2/docs"

	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

var (
	__PORT__ = ":9090"
)

func main() {

	db := database.InitDB()
	controllerOrders := &controller.NewDB{DB: db}
	router := gin.Default()

	//GET
	router.GET("/orders", controllerOrders.GetOrdered)

	//POST
	router.POST("/orders", controllerOrders.CreateOrder)

	//PUT
	router.PUT("/orders/:orderId", controllerOrders.UpdateOrdered)

	//DELETE
	router.DELETE("/orders/:orderId", controllerOrders.DeletedOrder)

	//swagg
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(__PORT__)
}
