package service

import (
	"github.com/3boku/backend1/repositories"
	"github.com/3boku/backend1/types"
	"net/http"
)

type ShopService struct {
	ShopRepository *repositories.ShopRepository
}

func (s *ShopService) CreateShop(shop types.ShopDAO) (statusCode int, data types.Shop, err error) {
	data, err = s.ShopRepository.Create(shop)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusCreated, data, nil
}

func (s *ShopService) UpdateShop(id string, shop *types.ShopDAO) (statusCode int, data *types.Shop, err error) {
	data, err = s.ShopRepository.Update(id, shop)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, data, nil
}

func (s *ShopService) DeleteShop(id string) (statusCode int, data string, err error) {
	err = s.ShopRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, "Deleted", nil
}

func (s *ShopService) GetShop() (statusCode int, data *[]types.Shop, err error) {
	data, err = s.ShopRepository.SelectALL()
	if err != nil {
		return http.StatusInternalServerError, data, err
	}
	return http.StatusOK, data, nil
}

func (s *ShopService) GetShopByID(id string) (statusCode int, data *types.Shop, err error) {
	data, err = s.ShopRepository.SelectByName(id)
	if err != nil {
		return http.StatusNotFound, data, err
	}
	return http.StatusOK, data, nil
}
