package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	SenderName      string `json:"sender_name"`
	SenderContact   string `json:"sender_contact"`
	SenderAddress   string `json:"sender_address"`
	ReceiverName    string `json:"receiver_name"`
	ReceiverContact string `json:"receiver_contact"`
	ReceiverAddress string `json:"receiver_address"`
	Item            string `json:"item"`
	Weight          string `json:"weight"`
	ItemAmount      string `json:"item_amount"`
	ShippingFee     string `json:"shipping_fee"`
}
