package repository

import (
	"api/internal/pkg/domain/domain_model/entity"
	"api/pkg/infrastucture/db"
	"fmt"
)

type productRepository struct {
	DB db.Database
}

func (u *productRepository) FirstProduct(condition entity.Product) (entity.Product, error) {
	product := entity.Product{}
	err := u.DB.First(&product, condition)
	return product, err
}
func (u *productRepository) FindProductList(condition entity.Product) (product []entity.Product, err error) {
	conditionString := fmt.Sprintf("deleted_at IS NULL AND name LIKE %s ", "'%"+condition.Name+"%'")
	if condition.BrandID > 0 {
		conditionString += fmt.Sprintf("AND brand_id=%d ", condition.BrandID)
	}
	if condition.CategoryID > 0 {
		conditionString += fmt.Sprintf("AND category_id=%d ", condition.CategoryID)
	}
	err = u.DB.Find(&product, conditionString)
	return
}
func (u *productRepository) FindProductListWithPagination(condition entity.Product, pageNum int, pageSize int) (product []entity.Product, err error) {
	offset := (pageNum - 1) * pageSize
	conditionString := fmt.Sprintf("deleted_at IS NULL AND name LIKE %s ", "'%"+condition.Name+"%'")
	if condition.BrandID > 0 {
		conditionString += fmt.Sprintf("AND brand_id=%d ", condition.BrandID)
	}
	if condition.CategoryID > 0 {
		conditionString += fmt.Sprintf("AND category_id=%d ", condition.CategoryID)
	}
	err = u.DB.FindWithPagination(&product, offset, pageSize, conditionString)
	return
}
func (u *productRepository) CreateProduct(product entity.Product) (entity.Product, error) {
	err := u.DB.Create(&product)
	return product, err
}
func (u *productRepository) DeleteProduct(product entity.Product) error {
	err := u.DB.Delete(&product)
	return err
}
func (u *productRepository) UpdateProduct(product, oldproduct entity.Product) (entity.Product, error) {
	err := u.DB.Update(&entity.Product{}, &oldproduct, &product)
	return product, err
}

func NewProductRepository(db db.Database) *productRepository {
	return &productRepository{
		DB: db,
	}
}
