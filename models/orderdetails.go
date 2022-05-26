package models

import (
	"log"
	"time"

	"github.com/dgquijote/be-screening/database"
)

type OrderDetails struct {
	Id              uint      `gorm:"primary_key" json:"id"`
	OrderId         int       `json:"order_id"`
	Order           Order     `json:"-"`
	TrackingNumber  string    `json:"tracking_number"`
	ShippingCourier string    `json:"shipping_courier"`
	ShippingStatus  string    `json:"shipping_status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func MigrateOrderDetails() {

	result := map[string]interface{}{}

	record := database.Instance.Model(&OrderDetails{}).First(&result)

	if record.Error != nil {
		database.Instance.AutoMigrate(&OrderDetails{})
		date_ordered, _ := time.Parse(time.RFC3339, "2022-05-01T08:34:45Z")
		date_pickedup, _ := time.Parse(time.RFC3339, "2022-05-02T09:12:02Z")
		date_DC, _ := time.Parse(time.RFC3339, "2022-05-02T22:08:50Z")
		date_delivery, _ := time.Parse(time.RFC3339, "2022-05-04T08:15:13Z")

		date_ordered2, _ := time.Parse(time.RFC3339, "2022-05-20T13:10:20Z")

		var records = []OrderDetails{
			{
				OrderId:         1,
				TrackingNumber:  "N/A",
				ShippingCourier: "N/A",
				ShippingStatus:  "Preparing Order",
				CreatedAt:       date_ordered,
				UpdatedAt:       date_ordered,
			},
			{
				OrderId:         1,
				TrackingNumber:  "794832149012",
				ShippingCourier: "FedEx",
				ShippingStatus:  "Picked-up by courier",
				CreatedAt:       date_pickedup,
				UpdatedAt:       date_pickedup,
			},
			{
				OrderId:         1,
				TrackingNumber:  "794832149012",
				ShippingCourier: "FedEx",
				ShippingStatus:  "Arrived on DC",
				CreatedAt:       date_DC,
				UpdatedAt:       date_DC,
			},
			{
				OrderId:         1,
				TrackingNumber:  "794832149012",
				ShippingCourier: "FedEx",
				ShippingStatus:  "Out for delivery",
				CreatedAt:       date_delivery,
				UpdatedAt:       date_delivery,
			},
			{
				OrderId:         2,
				TrackingNumber:  "N/A",
				ShippingCourier: "N/A",
				ShippingStatus:  "Preparing Order",
				CreatedAt:       date_ordered2,
				UpdatedAt:       date_ordered2,
			},
		}

		database.Instance.Create(&records)

		log.Println("Database Order Details Migration Completed!")
	}
}

func GetOrderTrackingDetails(id uint) []OrderDetails {
	var details []OrderDetails
	database.Instance.Where("order_id = ?", id).Find(&details)
	return details
}

func AddDetails(o Order) {
	item := OrderDetails{
		OrderId:         int(o.Id),
		TrackingNumber:  "N/A",
		ShippingCourier: "N/A",
		ShippingStatus:  "Preparing Order",
		CreatedAt:       o.CreatedAt,
		UpdatedAt:       o.UpdatedAt,
	}

	database.Instance.Omit("Seller").Omit("Receiver").Create(&item)
}
