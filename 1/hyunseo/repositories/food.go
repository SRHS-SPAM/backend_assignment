package repositories

import (
	"github.com/3boku/backend1/types"
	"gorm.io/gorm"
)

type FoodRepository struct {
	DB *gorm.DB
}

func (s *FoodRepository) SelectALL() (data *[]types.Shop, err error) {
	err = s.DB.Find(&data).Error
	return data, err
}

func (s *FoodRepository) SelectByName(name string) (data *types.Shop, err error) {
	err = s.DB.First(&data, "name =?", name).Error
	return data, err
}

func (s *FoodRepository) Create(shop *types.ShopDAO) (data types.Shop, err error) {
	input := types.Shop{
		Name:      shop.Name,
		OwnerName: shop.OwnerName,
		Category:  shop.Category,
	}
	err = s.DB.Create(&input).Error
	return input, err
}

func (s *FoodRepository) Update(id string, shop *types.ShopDAO) (data *types.Shop, err error) {
	input := types.Shop{
		Name:      shop.Name,
		OwnerName: shop.OwnerName,
		Category:  shop.Category,
	}
	err = s.DB.Model(&types.Shop{}).Where("id =?", id).Save(input).Error
	return data, err
}

func (s *FoodRepository) Delete(id string) (err error) {
	err = s.DB.Delete(&types.Shop{}, id).Error
	return err
}
