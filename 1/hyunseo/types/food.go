package types

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	ShopID string `json:"shopID"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Like   int    `json:"like"`
}

type FoodDAO struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Like  int    `json:"like"`
}
