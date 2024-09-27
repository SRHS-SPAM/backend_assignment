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
	err = f.DB.Model(&types.Food{}).Where("id = ?", id).Updates(input).Error
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

func (f *FoodRepository) Like(id string) (data *types.Food, err error) {
	food, err := f.SelectByName(id)
	if err != nil {
		return nil, err
	}

	input := &types.Food{
		Like: food.Like + 1,
	}

	err = f.DB.Model(&types.Food{}).Where("id = ?", id).Updates(input).Error
	if err != nil {
		return nil, err
	}

	// 업데이트된 데이터를 반환
	err = f.DB.First(&data, id).Error
	return data, err
}

func (f *FoodRepository) UnLike(id string) (data *types.Food, err error) {
	food, err := f.SelectByName(id)
	if err != nil {
		return nil, err
	}

	if food.Like > 0 {
		food.Like -= 1 // 1일 때 0으로, 그 이상의 경우에는 1씩 감소
	}
	if food.Like == 0 {
		food.Like = 0
	}

	input := &types.Food{
		Like: food.Like,
	}

	err = f.DB.Model(&types.Food{}).Where("id = ?", id).Updates(input).Error
	if err != nil {
		return nil, err
	}

	// 업데이트된 데이터를 반환
	err = f.DB.First(&data, id).Error
	return data, err
}
