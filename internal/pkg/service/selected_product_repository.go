package service

import "api/internal/pkg/domain/domain_model/entity"

type SelectedProductRepository interface {
	FindSelectedProductList(condition entity.SelectedProduct) ([]entity.SelectedProduct, error)
	CreateSelectedProduct(selectedProduct entity.SelectedProduct) (entity.SelectedProduct, error)
	DeleteSelectedProduct(selectedProduct entity.SelectedProduct) error
	UpdateSelectedProduct(newSelectedProduct, oldSelectedProduct entity.SelectedProduct) (entity.SelectedProduct, error)
}
