package repository

import (
	"api/internal/pkg/domain/domain_model/entity"
	"api/pkg/infrastucture/db"
)

type brandRepository struct {
	DB db.Database
}

func (u *brandRepository) FindbrandList(condition entity.Brand) (brand []entity.Brand, err error) {
	err = u.DB.Find(&brand, condition)
	return
}
func (u *brandRepository) Createbrand(brand entity.Brand) (entity.Brand, error) {
	err := u.DB.Create(&brand)
	return brand, err
}
func (u *brandRepository) Deletebrand(brand entity.Brand) error {
	err := u.DB.Delete(&brand)
	return err
}
func (u *brandRepository) Updatebrand(brand, oldbrand entity.Brand) (entity.Brand, error) {
	err := u.DB.Update(&entity.Brand{}, &oldbrand, &brand)
	return brand, err
}

func NewbrandRepository(db db.Database) *brandRepository {
	return &brandRepository{
		DB: db,
	}
}
