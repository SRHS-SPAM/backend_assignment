package repositories

import (
	"errors"
	"github.com/3boku/backend1/types"
	"gorm.io/gorm"
)

type FoodRepository struct {
	DB *gorm.DB
}

func (f *FoodRepository) SelectALL(shopID string) (data *[]types.Food, err error) {
	err = f.DB.Find(&data, "shop_id =?", shopID).Error
	return data, err
}

func (f *FoodRepository) SelectByName(id string) (data *types.Food, err error) {
	err = f.DB.First(&data, "id =?", id).Error
	if data.ID == 0 {
		return nil, errors.New("food not found")
	}
	return data, err
}

func (f *FoodRepository) Create(shopID string, food *types.FoodInput) (data *types.FoodInput, err error) {
	input := &types.Food{
		Name:   food.Name,
		Price:  food.Price,
		Like:   0,
		ShopID: shopID,
	}
	err = f.DB.Create(input).Error
	return food, err
}

func (f *FoodRepository) Update(id string, food *types.FoodInput) (data *types.Food, err error) {
	input := &types.Food{
		Name:  food.Name,
		Price: food.Price,
	}

	// Updates를 사용하여 지정된 필드만 업데이트
	err = f.DB.Model(&types.Shop{}).Where("id = ?", id).Updates(input).Error
	if err != nil {
		return nil, err
	}

	// 업데이트된 데이터를 반환
	err = f.DB.First(&data, id).Error
	return data, err
}

func (f *FoodRepository) Delete(id string) (err error) {
	err = f.DB.Delete(&types.Food{}, id).Error
	return err
}
