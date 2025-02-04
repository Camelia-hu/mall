package module

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Uid       uint `json:"uid"`
	ProductId uint `json:"productId"`
	Quantity  int  `json:"quantity"`
}
