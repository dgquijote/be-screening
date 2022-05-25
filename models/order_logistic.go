package models

import "gorm.io/gorm"

type OrderLogistics struct {
	gorm.Model
	OrderId         string `json:"order_id"`
	TrackingNumber  string `json:"tracking_number"`
	ShippingCourier string `json:"shipping_courier"`
	ShippingStatus  string `json:"shipping_status"`
}
