package handler

import (
	"api/internal/pkg/domain/domain_model/dto"
	"api/internal/pkg/usecase"
	"api/pkg/infrastucture/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(db db.Database) *CategoryHandler {
	u := usecase.NewCategoryUsecase(db)
	return &CategoryHandler{
		categoryUsecase: u,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	req := dto.CreateCategoryRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	res, err := h.categoryUsecase.CreateCategory(req)
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

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	req := dto.UpdateCategoryRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	res, err := h.categoryUsecase.UpdateCategory(req)
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
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	req := dto.DeleteCategoryRequest{}
	err := c.ShouldBind(req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	err = h.categoryUsecase.DeleteCategory(req.CategoryID)
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
func (h *CategoryHandler) GetCategoryList(c *gin.Context) {
	res, err := h.categoryUsecase.GetCategoryList()
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
