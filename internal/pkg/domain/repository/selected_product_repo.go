package repository

import (
	"api/internal/pkg/domain/domain_model/entity"
	"api/pkg/infrastucture/db"
)

type selectedProductRepository struct {
	DB db.Database
}

func (u *selectedProductRepository) FindSelectedProductList(condition entity.SelectedProduct) (selectedProduct []entity.SelectedProduct, err error) {
	err = u.DB.Find(&selectedProduct, condition)
	return
}
func (u *selectedProductRepository) CreateSelectedProduct(selectedProduct entity.SelectedProduct) (entity.SelectedProduct, error) {
	err := u.DB.Create(&selectedProduct)
	return selectedProduct, err
}
func (u *selectedProductRepository) DeleteSelectedProduct(selectedProduct entity.SelectedProduct) error {
	err := u.DB.Delete(&selectedProduct)
	return err
}
func (u *selectedProductRepository) UpdateSelectedProduct(selectedProduct, oldselectedProduct entity.SelectedProduct) (entity.SelectedProduct, error) {
	err := u.DB.Update(&entity.SelectedProduct{}, &oldselectedProduct, &selectedProduct)
	return selectedProduct, err
}

func NewSelectedProductRepository(db db.Database) *selectedProductRepository {
	return &selectedProductRepository{
		DB: db,
	}
}
