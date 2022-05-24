package models

import "gorm.io/gorm"

type OrderLogistics struct {
	gorm.Model
	OrderId        string `json:"order_id"`
	ShippingStatus string `json:"shipping_status"`
}
