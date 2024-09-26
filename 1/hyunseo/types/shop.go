package types

import (
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Name      string `json:"name"`
	OwnerName string `json:"owner_name"`
	Category  string `json:"category"`
}

type ShopDAO struct {
	Name      string `json:"name"`
	OwnerName string `json:"owner_name"`
	Category  string `json:"category"`
}
