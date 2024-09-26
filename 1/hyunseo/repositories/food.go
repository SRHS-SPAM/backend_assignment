package repositories

//import (
//	"github.com/3boku/backend1/types"
//	"github.com/google/uuid"
//	"gorm.io/gorm"
//)
//
//type FoodRepository struct {
//	DB *gorm.DB
//}
//
//func (s *FoodRepository) SelectALL() (data *[]types.FoodDAO, err error) {
//	err = s.DB.Find(&data).Error
//	return data, err
//}
//
//func (s *FoodRepository) SelectByName(name string) (data *types.FoodDAO, err error) {
//	err = s.DB.First(&data, "name =?", name).Error
//	return data, err
//}
//
//func (s *FoodRepository) Create(shop *types.FoodDAO) (data *types.FoodDAO, err error) {
//	id, _ := uuid.NewUUID()
//	input := types.Shop{
//		UUID:      id,
//		Name:      shop.Name,
//		OwnerName: shop.OwnerName,
//		Category:  shop.Category,
//	}
//	err = s.DB.Create(&input).Error
//	return shop, err
//}
//
//func (s *FoodRepository) Update(id string, shop *types.ShopDAO) (data *types.ShopDAO, err error) {
//	input := types.Shop{
//		Name:      shop.Name,
//		OwnerName: shop.OwnerName,
//		Category:  shop.Category,
//	}
//	err = s.DB.Model(&types.Shop{}).Where("id =?", id).Save(input).Error
//	return data, err
//}
//
//func (s *FoodRepository) Delete(id string) (err error) {
//	err = s.DB.Delete(&types.FoodDAO{}, id).Error
//	return err
//}
