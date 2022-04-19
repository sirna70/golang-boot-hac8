package controller

import (
	"time"

	"gorm.io/gorm"
)

//Item represents the model for any item
type Items struct {
	gorm.Model
	Order_Id    uint   `json:"order_id"`
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

//order represented the model for any order
type Orders struct {
	gorm.Model
	CustomerName string    `json:"customer_name"`
	Ordered_At   time.Time `json:"Ordered_At"`
	Items        []Items   `gorm:"foreignKey:Order_Id"`
}
