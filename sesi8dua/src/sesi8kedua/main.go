package main

import (
	"fmt"
	"sesi8-project/configs"
	"sesi8-project/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.NewPostgres()
	if db == nil {
		fmt.Println("running db is fail")
	}

	personController := controllers.NewControllerPerson(db)

	router := gin.Default()

	router.GET("/persons", personController.GetPeople)
	router.POST("/persons", personController.CreatePerson)

	router.GET("/persons/:id", personController.FindPersonByID)
	router.PUT("/persons/:id", personController.UpdatePersonByID)
	router.DELETE("/persons/:id", personController.DeletePersonByID)

	router.Run(":8889")
}
