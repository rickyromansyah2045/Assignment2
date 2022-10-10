package models

import "time"

type Order struct {
	ID           int    `gorm:"primaryKey" json:"order_id"`
	CustomerName string `gorm:"not null; type:varchar(50)" json:"customer_name"`
	Items        []Item
	OrderedAt    time.Time
}
