package service

import "api/internal/pkg/domain/domain_model/entity"

type CategoryRepository interface {
	FindCategoryList(condition entity.Category) ([]entity.Category, error)
	CreateCategory(category entity.Category) (entity.Category, error)
	DeleteCategory(category entity.Category) error
	UpdateCategory(category, oldcategory entity.Category) (entity.Category, error)
}
