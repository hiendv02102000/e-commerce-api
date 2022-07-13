package service

import "api/internal/pkg/domain/domain_model/entity"

type ProductRepository interface {
	FirstProduct(condition entity.Product) (entity.Product, error)
	FindProductList(condition entity.Product) ([]entity.Product, error)
	FindProductListWithPagination(condition entity.Product, pageNum int, pageSize int) ([]entity.Product, error)
	CreateProduct(product entity.Product) (entity.Product, error)
	DeleteProduct(product entity.Product) error
	UpdateProduct(newProduct, oldProduct entity.Product) (entity.Product, error)
}
