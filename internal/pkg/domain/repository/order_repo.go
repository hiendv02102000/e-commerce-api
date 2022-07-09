package repository

import (
	"api/internal/pkg/domain/domain_model/entity"
	"api/pkg/infrastucture/db"
)

type orderRepository struct {
	DB db.Database
}

func (u *orderRepository) FirstOrder(condition entity.Order) (entity.Order, error) {
	order := entity.Order{}
	err := u.DB.First(&order, condition)
	return order, err
}
func (u *orderRepository) FindOrderList(condition entity.Order) (order []entity.Order, err error) {
	err = u.DB.Find(&order, condition)
	return
}
func (u *orderRepository) FindOrderListWithPagination(condition entity.Order, pageNum int, pageSize int) (order []entity.Order, err error) {
	offset := (pageNum - 1) * pageSize
	err = u.DB.FindWithPagination(&order, offset, pageSize, condition)
	return
}
func (u *orderRepository) CreateOrder(order entity.Order) (entity.Order, error) {
	err := u.DB.Create(&order)
	return order, err
}
func (u *orderRepository) DeleteOrder(order entity.Order) error {
	err := u.DB.Delete(&order)
	return err
}
func (u *orderRepository) UpdateOrder(order, oldorder entity.Order) (entity.Order, error) {
	err := u.DB.Update(entity.Order{}, &oldorder, &order)
	return order, err
}

func NeworderRepository(db db.Database) *orderRepository {
	return &orderRepository{
		DB: db,
	}
}
