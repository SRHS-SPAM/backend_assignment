package types

import "gorm.io/gorm"

type FoodDTO struct {
	gorm.Model
	Name  string `json:"name"`
	Price int    `json:"price"`
	Like  int    `json:"like"`
}

type FoodDAO struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Like  int    `json:"like"`
}
