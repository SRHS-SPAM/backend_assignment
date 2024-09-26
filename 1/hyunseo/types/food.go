package types

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name  string `json:"name"`
	Price int    `json:"price"`
	Like  int    `json:"like"`
}
