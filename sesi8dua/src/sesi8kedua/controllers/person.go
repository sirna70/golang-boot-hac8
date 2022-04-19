package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sesi8-project/models"
	"sesi8-project/views"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type controllerPerson struct {
	db *gorm.DB
}

func NewControllerPerson(db *gorm.DB) *controllerPerson {
	return &controllerPerson{
		db: db,
	}
}

func (in *controllerPerson) GetPeople(c *gin.Context) {
	var (
		people []models.Person
		result gin.H
	)

	in.db.Find(&people)

	if len(people) == 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		var personViews []views.PersonView

		for _, p := range people {
			personView := views.PersonView{
				ID:        int(p.ID),
				FirstName: p.FirstName,
				LastName:  p.LastName,
			}

			personViews = append(personViews, personView)
		}

		result = gin.H{
			"result": personViews,
			"count":  len(personViews),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (in *controllerPerson) CreatePerson(c *gin.Context) {
	var person models.Person

	err := json.NewDecoder(c.Request.Body).Decode(&person)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Create(&person).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": person,
	})

}

func (in *controllerPerson) DeletePersonByID(c *gin.Context) {
	var (
		person models.Person
	)

	id := c.Param("id")

	err := in.db.First(&person, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Delete(&person).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "delete data success !",
	})
}

func (in *controllerPerson) UpdatePersonByID(c *gin.Context) {
	var (
		person    models.Person
		newPerson models.Person
	)

	id := c.Param("id")

	err := in.db.First(&person, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&newPerson)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = in.db.Model(&person).Updates(newPerson).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "update data success !",
		"data": newPerson,
	})
}
func (in *controllerPerson) FindPersonByID(c *gin.Context) {
	var (
		person models.Person
	)

	id := c.Param("id")

	err := in.db.First(&person, id).Error
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

		}
		return
	}

	view := views.PersonName{
		FirstName: person.FirstName,
		LastName:  person.LastName,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": view,
	})
}
