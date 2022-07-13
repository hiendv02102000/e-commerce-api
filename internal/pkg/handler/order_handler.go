package handler

import (
	"api/internal/pkg/domain/domain_model/dto"
	"api/internal/pkg/usecase"
	"api/pkg/infrastucture/db"
	"api/pkg/share/middleware"
	"net/http"
	"strings"

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

func (h *OrderHandler) GetOrderListWithAdmin(c *gin.Context) {
	req := dto.GetOrderListRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return

	}
	res, err := h.OrderUsecase.GetOrderList(dto.GetOrderListRequest{})
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	c.JSON(http.StatusOK, data)
}
func (h *OrderHandler) GetOrderListWithCustomer(c *gin.Context) {
	req := dto.GetOrderListRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	user := middleware.GetUserFromContext(c)
	req.CustomerID = user.ID
	res, err := h.OrderUsecase.GetOrderList(req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	c.JSON(http.StatusOK, data)
}
func (h *OrderHandler) GetOrderInfoWithAdmin(c *gin.Context) {

	transactionCode := c.Param("transaction_code")
	if len(transactionCode) == 0 {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  "transaction_code is required",
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	res, err := h.OrderUsecase.GetOrderInfo(transactionCode, 0)
	if err != nil {
		code := http.StatusBadRequest
		if strings.Contains(err.Error(), "method not allowed") {
			code = http.StatusMethodNotAllowed
		}
		data := dto.BaseResponse{
			Status: code,
			Error:  err.Error(),
		}
		c.JSON(code, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	c.JSON(http.StatusOK, data)
}
func (h *OrderHandler) GetOrderInfoWithCustomer(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	transactionCode := c.Param("transaction_code")
	if len(transactionCode) == 0 {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  "transaction_code is required",
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	res, err := h.OrderUsecase.GetOrderInfo(transactionCode, user.ID)
	if err != nil {
		code := http.StatusBadRequest
		if strings.Contains(err.Error(), "method not allowed") {
			code = http.StatusMethodNotAllowed
		}
		data := dto.BaseResponse{
			Status: code,
			Error:  err.Error(),
		}
		c.JSON(code, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	c.JSON(http.StatusOK, data)
}
func (h *OrderHandler) ChangeStatusOrder(c *gin.Context) {
	req := dto.UpdateStatucOrderRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}

	res, err := h.OrderUsecase.ChangeStatusOrder(req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	c.JSON(http.StatusOK, data)
}
func (h *OrderHandler) GetCart(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	res, err := h.OrderUsecase.GetCart(user)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	c.JSON(http.StatusOK, data)
}
func (h *OrderHandler) CreateCart(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	res, err := h.OrderUsecase.CreateCart(user)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusCreated,
		Result: res,
	}
	c.JSON(http.StatusCreated, data)
}
func (h *OrderHandler) PutProductToCart(c *gin.Context) {
	req := dto.PutProductToCartRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	err = h.OrderUsecase.PutProductToCart(req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusNoContent,
	}
	c.JSON(http.StatusNoContent, data)
}
func (h *OrderHandler) DeleteProductFromCart(c *gin.Context) {
	req := dto.DeleteProductFromCartRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	user := middleware.GetUserFromContext(c)
	req.CustomerID = user.ID
	err = h.OrderUsecase.DeleteProductFromCart(req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusNoContent,
	}
	c.JSON(http.StatusNoContent, data)
}
func (h *OrderHandler) CancelOrder(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	req := dto.CancelOrderRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	err = h.OrderUsecase.CancelOrder(req.TransactionCode, user.ID)
	if err != nil {
		code := http.StatusBadRequest
		if strings.Contains(err.Error(), "method not allowed") {
			code = http.StatusMethodNotAllowed
		}
		data := dto.BaseResponse{
			Status: code,
			Error:  err.Error(),
		}
		c.JSON(code, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: "success",
	}
	c.JSON(http.StatusOK, data)
}
func (h *OrderHandler) ConfirmOrder(c *gin.Context) {
	user := middleware.GetUserFromContext(c)
	err := h.OrderUsecase.ConfirmOrder(user)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: "success",
	}
	c.JSON(http.StatusOK, data)
}
