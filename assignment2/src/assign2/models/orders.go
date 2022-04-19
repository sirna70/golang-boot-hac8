package models

import (
	"time"

	"gorm.io/gorm"
)

// order represented the model for any order
type Orders struct {
	gorm.Model
	CustomerName string    `json:"customer_name"`
	Ordered_At   time.Time `json:"ordered_at"`
	Items        []Items   `gorm:"foreignKey:Order_Id"`
}
