package repository

import (
	"api/internal/pkg/domain/domain_model/entity"
	"api/pkg/infrastucture/db"
)

type userRepository struct {
	DB db.Database
}

func (u *userRepository) FirstUser(condition entity.Users) (entity.Users, error) {
	user := entity.Users{}
	err := u.DB.First(&user, condition)
	return user, err
}
func (u *userRepository) FindUserList(condition entity.Users) (user []entity.Users, err error) {
	err = u.DB.Find(&user, condition)
	return
}

func (u *userRepository) CreateUser(user entity.Users) (entity.Users, error) {
	err := u.DB.Create(&user)
	return user, err
}
func (u *userRepository) DeleteUser(user entity.Users) error {
	err := u.DB.Delete(&user)
	return err
}
func (u *userRepository) UpdateUser(user, oldUser entity.Users) (entity.Users, error) {
	err := u.DB.Update(entity.Users{}, &oldUser, &user)
	return user, err
}

func NewUserRepository(db db.Database) *userRepository {
	return &userRepository{
		DB: db,
	}
}
