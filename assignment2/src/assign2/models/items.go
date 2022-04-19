package models

import (
	"gorm.io/gorm"
)

// Item represents the model for any item
type Items struct {
	gorm.Model
	Order_Id    uint   `json:"order_id"`
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
