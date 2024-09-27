package repositories

import (
	"github.com/3boku/backend1/types"
	"gorm.io/gorm"
)

type ShopRepository struct {
	DB *gorm.DB
}

func (s *ShopRepository) SelectALL() (data *[]types.Shop, err error) {
	err = s.DB.Find(&data).Error
	return data, err
}

func (s *ShopRepository) SelectByID(id string) (data *types.Shop, err error) {
	err = s.DB.First(&data, "id =?", id).Error
	return data, err
}

func (s *ShopRepository) Create(shop types.ShopDAO) (data types.Shop, err error) {
	input := types.Shop{
		Name:      shop.Name,
		OwnerName: shop.OwnerName,
		Category:  shop.Category,
	}
	err = s.DB.Create(&input).Error
	return input, err
}

func (s *ShopRepository) Update(id string, shop *types.ShopDAO) (data *types.Shop, err error) {
	input := &types.Shop{
		Name:      shop.Name,
		OwnerName: shop.OwnerName,
		Category:  shop.Category,
	}

	// Updates를 사용하여 지정된 필드만 업데이트
	err = s.DB.Model(&types.Shop{}).Where("id = ?", id).Updates(input).Error
	if err != nil {
		return nil, err
	}

	// 업데이트된 데이터를 반환
	err = s.DB.First(&data, id).Error
	return data, err
}

func (s *ShopRepository) Delete(id string) (err error) {
	err = s.DB.Delete(&types.Shop{}, id).Error
	return err
}
