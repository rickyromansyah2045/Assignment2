package controllers

import (
	"assignment-2/helpers"
	"assignment-2/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	var newReqOrder ReqOrder

	err := ctx.ShouldBindJSON(&newReqOrder)
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	newOrder := models.Order{
		CustomerName: newReqOrder.CustomerName,
	}

	err = o.db.Create(&newOrder).Error
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	var newResItem []ResItem

	for _, item := range newReqOrder.Items {
		newItem := models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     newOrder.ID,
		}

		err = o.db.Create(&newItem).Error
		if err != nil {
			fmt.Println("Error create Order")
			helpers.BadRequestResponse(ctx, err.Error())
			return
		}

		newResItem = append(newResItem, ResItem{
			ID:          newItem.ID,
			ItemCode:    newItem.ItemCode,
			Description: newItem.Description,
			Quantity:    newItem.Quantity,
		})
	}

	newResOrder := ResOrder{
		ID:           newOrder.ID,
		CustomerName: newOrder.CustomerName,
		OrderedAt:    *newOrder.OrderedAt,
		Items:        newResItem,
	}

	helpers.WriteJsonResponse(ctx, http.StatusCreated, gin.H{
		"status": true,
		"data":   newResOrder,
	})
}

func (o *OrderController) GetOrders(ctx *gin.Context) {

}

func (o *OrderController) UpdatedOrder(ctx *gin.Context) {

}

func (o *OrderController) DeleteOrder(ctx *gin.Context) {

}
