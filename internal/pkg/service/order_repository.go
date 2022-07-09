package service

import "api/internal/pkg/domain/domain_model/entity"

type OrderRepository interface {
	FirstOrder(condition entity.Order) (entity.Order, error)
	FindOrderList(condition entity.Order) ([]entity.Order, error)
	FindOrderListWithPagination(condition entity.Order, pageNum int, pageSize int) ([]entity.Order, error)
	CreateOrder(order entity.Order) (entity.Order, error)
	DeleteOrder(order entity.Order) error
	UpdateOrder(newOrder, oldOrder entity.Order) (entity.Order, error)
}
