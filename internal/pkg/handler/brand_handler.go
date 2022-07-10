package handler

import (
	"api/internal/pkg/domain/domain_model/dto"
	"api/internal/pkg/usecase"
	"api/pkg/infrastucture/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	brandUsecase usecase.BrandUsecase
}

func NewBrandHandler(db db.Database) *BrandHandler {
	u := usecase.NewBrandUsecase(db)
	return &BrandHandler{
		brandUsecase: u,
	}
}

func (h *BrandHandler) CreateBrand(c *gin.Context) {
	req := dto.CreateBrandRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	res, err := h.brandUsecase.CreateBrand(req)
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
	c.JSON(http.StatusCreated, data)
}

func (h *BrandHandler) UpdateBrand(c *gin.Context) {
	req := dto.UpdateBrandRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	res, err := h.brandUsecase.UpdateBrand(req)
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
func (h *BrandHandler) DeleteBrand(c *gin.Context) {
	req := dto.DeleteBrandRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	err = h.brandUsecase.DeleteBrand(req.BrandID)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
func (h *BrandHandler) GetBrandList(c *gin.Context) {
	res, err := h.brandUsecase.GetBrandList()
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
