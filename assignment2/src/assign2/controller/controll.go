package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewDB struct {
	DB *gorm.DB
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order to database
// @Tags  createorders customer
// @Accept json
// @Produce json
// @Param order body Orders true "Create Order with payload "
// @Success 200 {object} Orders
// @Router /orders [post]
func (in *NewDB) CreateOrder(c *gin.Context) {
	var (
		order  Orders
		result gin.H
	)
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&order)
	if err != nil {
		result = gin.H{
			"status":  "Error",
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	err = in.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"status":  "Error",
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	err = in.DB.Model(&order).Association("Items").Append(&order.Items)
	if err != nil {
		result = gin.H{
			"status":  "Error",
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	in.DB.Preload("Items").First(&order, order.ID)
	result = gin.H{
		"status": "Succesfully",
		"code":   200,
		"result": order,
	}
	c.JSON(http.StatusOK, result)
}

// GetOrders godoc
// @Summary Get All Orders
// @Description Get all orders from database
// @Tags  orders
// @Accept json
// @Produce json
// @Param order body Orders true "Get Orders"
// @Success 200 {object} Orders
// @Router /orders [get]
func (in *NewDB) GetOrdered(c *gin.Context) {
	var (
		orders []Orders
		result gin.H
	)
	err := in.DB.Preload("Items").Find(&orders).Error
	if err != nil {
		result = gin.H{
			"status":  "Error",
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	} else {
		result = gin.H{
			"status": "Succesfully",
			"code":   200,
			"result": orders,
			"count":  len(orders),
		}
		c.JSON(http.StatusOK, result)
		return
	}
}

// UpdateOrder godoc
// @Summary Create a new order
// @Description Update exists order from database
// @Tags  createorders customer
// @Accept json
// @Produce json
// @Param order body Orders true "Update Order with payload "
// @Success 200 {object} Orders
// @Router /orders/:orderId [put]
func (in *NewDB) UpdateOrdered(c *gin.Context) {
	var (
		order       Orders
		updateOrder Orders
		result      gin.H
	)
	orderID := c.Param("orderId")
	decoder := json.NewDecoder(c.Request.Body)
	err := in.DB.First(&order, orderID).Error
	if err != nil {
		result = gin.H{
			"status":  "Error",
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	err = decoder.Decode(&updateOrder)
	if err != nil {
		result = gin.H{
			"status":  "Error",
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	err = in.DB.First(&order, orderID).Updates(&updateOrder).Error
	if err != nil {
		result = gin.H{
			"status":  "Error",
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	for _, item := range updateOrder.Items {
		if item.ID == 0 {
			err = in.DB.Create(&item).Error
			if err != nil {
				result = gin.H{
					"status":  "Error",
					"code":    http.StatusBadRequest,
					"message": err.Error(),
				}
				c.JSON(http.StatusBadRequest, result)
				return
			}
		} else {
			err = in.DB.Model(&item).Where("id = ?", item.ID).Updates(&item).Error
			if err != nil {
				result = gin.H{
					"status":  "Error",
					"code":    http.StatusBadRequest,
					"message": err.Error(),
				}
				c.JSON(http.StatusBadRequest, result)
				return
			}
		}
	}
	err = in.DB.Preload("Items").First(&order, orderID).Error
	if err != nil {
		result = gin.H{
			"status":  "Error",
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result = gin.H{
		"status": "Succesfully",
		"code":   200,
		"result": order,
	}
	c.JSON(http.StatusOK, result)
}

// DeleteOrder godoc
// @Summary Delete order with associations
// @Description Delete order with associations
// @Tags  deleteorders customer
// @Accept json
// @Produce json
// @Param order body Orders true "Delete by param "
// @Success 200 {object} Orders
// @Router /orders/:orderId [delete]
func (in *NewDB) DeletedOrder(c *gin.Context) {
	var (
		order  Orders
		items  []Items
		result gin.H
	)
	orderID := c.Param("orderId")
	err := in.DB.Preload("Items").First(&order, orderID).Error
	if err != nil {
		result = gin.H{
			"status":  "Error",
			"code":    http.StatusNotFound,
			"message": err.Error(),
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	// delete all items from orders
	in.DB.Model(&order).Association("Items").Find(&items)
	in.DB.Model(&order).Association("Items").Delete(&order.Items)
	for _, item := range items {
		err = in.DB.Delete(&item).Error
		if err != nil {
			result = gin.H{
				"status":  "Error",
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			}
			c.JSON(http.StatusBadRequest, result)
			return
		}
	}
	err = in.DB.Delete(&order).Error
	if err != nil {
		result = gin.H{
			"status":  "Error",
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result = gin.H{
		"status":  "Succesfully",
		"code":    200,
		"message": "Order deleted succesfully",
	}
	c.JSON(200, result)
}
