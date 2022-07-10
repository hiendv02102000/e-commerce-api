package repository

import (
	"api/internal/pkg/domain/domain_model/entity"
	"api/pkg/infrastucture/db"
)

type categoryRepository struct {
	DB db.Database
}

func (u *categoryRepository) FindCategoryList(condition entity.Category) (category []entity.Category, err error) {
	err = u.DB.Find(&category, condition)
	return
}
func (u *categoryRepository) CreateCategory(category entity.Category) (entity.Category, error) {
	err := u.DB.Create(&category)
	return category, err
}
func (u *categoryRepository) DeleteCategory(category entity.Category) error {
	err := u.DB.Delete(&category)
	return err
}
func (u *categoryRepository) UpdateCategory(category, oldcategory entity.Category) (entity.Category, error) {
	err := u.DB.Update(&entity.Category{}, &oldcategory, &category)
	return category, err
}

func NewCategoryRepository(db db.Database) *categoryRepository {
	return &categoryRepository{
		DB: db,
	}
}
