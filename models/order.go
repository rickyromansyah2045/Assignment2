package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           int        `gorm:"primaryKey" json:"id"`
	CustomerName string     `gorm:"not null;type:varchar(150)" json:"customer_name" valid:"required~customer name is required"`
	OrderedAt    *time.Time `gorm:"autoCreateTime" json:"ordered_at"`
	Items        []Item     `json:"items"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    gorm.DeletedAt
}
