package service

import "api/internal/pkg/domain/domain_model/entity"

type UserRepository interface {
	FirstUser(condition entity.Users) (entity.Users, error)
	FindUserList(condition entity.Users) (user []entity.Users, err error)
	CreateUser(user entity.Users) (entity.Users, error)
	DeleteUser(user entity.Users) error
	UpdateUser(newUser, oldUser entity.Users) (entity.Users, error)
}
