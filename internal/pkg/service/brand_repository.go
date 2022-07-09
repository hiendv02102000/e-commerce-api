package service

import "api/internal/pkg/domain/domain_model/entity"

type BrandRepository interface {
	FindBrandList(condition entity.Brand) ([]entity.Brand, error)
	CreateBrand(brand entity.Brand) (entity.Brand, error)
	DeleteBrand(brand entity.Brand) error
	UpdateBrand(newBrand, oldBrand entity.Users) (entity.Users, error)
}
