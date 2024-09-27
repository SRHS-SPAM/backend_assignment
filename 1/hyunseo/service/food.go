package service

import (
	"github.com/3boku/backend1/repositories"
	"github.com/3boku/backend1/types"
	"net/http"
)

type FoodService struct {
	FoodRepository *repositories.FoodRepository
}

func (f *FoodService) CreateFood(shopID string, food *types.FoodInput) (statusCode int, data *types.FoodInput, err error) {
	data, err = f.FoodRepository.Create(shopID, food)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusCreated, data, nil
}

func (f *FoodService) UpdateFood(id string, food *types.FoodInput) (statusCode int, data *types.Food, err error) {
	data, err = f.FoodRepository.Update(id, food)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, data, nil
}

func (f *FoodService) DeleteFood(id string) (statusCode int, data string, err error) {
	err = f.FoodRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, "Deleted", nil
}

func (f *FoodService) SeleteAll(id string) (statusCode int, data *[]types.Food, err error) {
	data, err = f.FoodRepository.SelectALL(id)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, data, nil
}

func (f *FoodService) SeleteByID(id string) (statusCode int, data *types.Food, err error) {
	data, err = f.FoodRepository.SelectByName(id)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, data, nil
}

func (f *FoodService) Like(id string) (statusCode int, data *types.Food, err error) {
	data, err = f.FoodRepository.Like(id)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, data, nil
}

func (f *FoodService) UnLike(id string) (statusCode int, data *types.Food, err error) {
	data, err = f.FoodRepository.UnLike(id)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, data, nil
}
