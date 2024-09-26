package repositories

import (
	"github.com/3boku/backend1/types"
	"gorm.io/gorm"
)

type ShopService struct {
	DB *gorm.DB
}

func (s *ShopService) SelectALL() (data []*types.Shop, err error) {
	err = s.DB.Find(&data).Error
	return data, err
}

func (s *ShopService) SelectByName(name string) (data *types.Shop, err error) {
	err = s.DB.First(&data, "name =?", name).Error
	return data, err
}

func (s *ShopService) Create(shop *types) (data *types.Shop, err error) {
	err = s.DB.Create(&shop).Error
	return data, err
}
