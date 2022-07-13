package handler

import (
	"api/internal/pkg/usecase"
	"api/pkg/infrastucture/db"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUsecase usecase.OrderUsecase
}

func NewOrderHandler(db db.Database) *OrderHandler {
	u := usecase.NewOrderUsecase(db)
	return &OrderHandler{
		OrderUsecase: u,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {

}

func (h *OrderHandler) GetOrderInfo(c *gin.Context) {

}
