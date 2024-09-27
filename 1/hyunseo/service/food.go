package service

import (
	"github.com/3boku/backend1/repositories"
	"github.com/3boku/backend1/types"
	"net/http"
)

type FoodService struct {
	foodRepository *repositories.FoodRepository
}

func (f *FoodService) CreateShop(shopID string, food *types.FoodInput) (statusCode int, data *types.FoodInput, err error) {
	data, err = f.foodRepository.Create(shopID, food)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusCreated, data, nil
}

func (f *FoodService) UpdateShop(id string, food *types.FoodInput) (statusCode int, data *types.Food, err error) {
	data, err = f.foodRepository.Update(id, food)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, data, nil
}

func (f *FoodService) DeleteShop(id string) (statusCode int, data string, err error) {
	err = f.foodRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, "Deleted", nil
}

func (f *FoodService) SeleteAll(id string) (statusCode int, data *[]types.Food, err error) {
	data, err = f.foodRepository.SelectALL(id)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, data, nil
}

func (f *FoodService) SeleteByID(id string) (statusCode int, data *types.Food, err error) {
	data, err = f.foodRepository.SelectByName(id)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, data, nil
}
