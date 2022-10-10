package models

type Item struct {
	ID          int    `gorm:"primaryKey" json:"item_id"`
	ItemCode    int    `json:"item_code"`
	Description string `gorm:"type:varchar(250)" json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int
}
