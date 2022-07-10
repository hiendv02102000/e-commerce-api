package repository

import (
	"api/internal/pkg/domain/domain_model/entity"
	"api/pkg/infrastucture/db"
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
	err = u.DB.Find(&product, condition)
	return
}
func (u *productRepository) FindProductListWithPagination(condition entity.Product, pageNum int, pageSize int) (product []entity.Product, err error) {
	offset := (pageNum - 1) * pageSize
	err = u.DB.FindWithPagination(&product, offset, pageSize, condition)
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

func NewproductRepository(db db.Database) *productRepository {
	return &productRepository{
		DB: db,
	}
}
