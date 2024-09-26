package service

import (
	"github.com/3boku/backend1/repositories"
	"github.com/3boku/backend1/types"
	"net/http"
)

type Shop struct {
	shopService *repositories.ShopService
}

func (s *Shop) CreateShop(shop *types.ShopDAO) (statusCode int, data *types.ShopDAO, err error) {
	data, err = s.shopService.Create(shop)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusCreated, data, nil
}

func (s *Shop) UpdateShop(id string, shop *types.ShopDAO) (statusCode int, data *types.ShopDAO, err error) {
	data, err = s.shopService.Update(id, shop)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, data, nil
}

func (s *Shop) DeleteShop(id string) (statusCode int, data string, err error) {
	err = s.shopService.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, data, err
	}

	return http.StatusOK, "Deleted", nil
}

func (s *Shop) GetShop() (statusCode int, data *[]types.ShopDAO, err error) {
	data, err = s.shopService.SelectALL()
	if err != nil {
		return http.StatusInternalServerError, data, err
	}
	return http.StatusOK, data, nil
}

func (s *Shop) GetShopByID(id string) (statusCode int, data *types.ShopDAO, err error) {
	data, err = s.shopService.SelectByName(id)
	if err != nil {
		return http.StatusNotFound, data, err
	}
	return http.StatusOK, data, nil
}
