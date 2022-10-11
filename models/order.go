package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Order struct {
	ID           int            `gorm:"primaryKey" json:"id"`
	CustomerName string         `gorm:"not null;type:varchar(150)" json:"customer_name" valid:"required~customer name is required"`
	OrderedAt    *time.Time     `gorm:"autoCreateTime" json:"ordered_at"`
	Items        []Item         `json:"items,omitempty"`
	CreatedAt    *time.Time     `json:"created_at,omitempty"`
	UpdatedAt    *time.Time     `json:"updated_at,omitempty"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (i *Order) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(i)
	if errCreate != nil {
		return errCreate
	}

	return
}
