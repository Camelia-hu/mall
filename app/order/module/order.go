package module

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Uid          uint32      `json:"uid"`
	UserCurrency string      `json:"userCurrency"`
	Address      string      `json:"address"`
	Email        string      `json:"email"`
	OrderItems   []OrderItem `gorm:"constraint:OnDelete:CASCADE;"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"orderID"`
	Cost      float32 `json:"cost"`
	ProductId uint32  `json:"productId"`
	Quantity  int32   `json:"quantity"`
	Order     Order   `gorm:"foreignKey:OrderID;references:ID"`
}
