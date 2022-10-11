package controllers

import (
	"time"

	"gorm.io/gorm"
)

type OrderController struct {
	db *gorm.DB
}

type ReqOrder struct {
	CustomerName string    `json:"customer_name"`
	Items        []ReqItem `json:"items"`
}

type ResOrder struct {
	ID           uint      `json:"id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []ResItem `json:"items,omitempty"`
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{
		db: db,
	}
}
