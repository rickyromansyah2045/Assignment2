package controllers

import (
	"gorm.io/gorm"
)

type ItemsController struct {
	db *gorm.DB
}

type ReqItem struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type ResItem struct {
	ID          uint   `json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func NewItemController(db *gorm.DB) *ItemsController {
	return &ItemsController{
		db: db,
	}
}
