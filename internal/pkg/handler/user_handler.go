package handler

import (
	"api/internal/pkg/domain/domain_model/dto"
	"api/internal/pkg/usecase"
	"api/pkg/infrastucture/db"
	"api/pkg/share/middleware"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewCustomerHandler(db db.Database) *UserHandler {
	u := usecase.NewCustomerUsecase(db)
	return &UserHandler{
		userUsecase: u,
	}
}
func NewAdminHandler(db db.Database) *UserHandler {
	u := usecase.NewAdminUsecase(db)
	return &UserHandler{
		userUsecase: u,
	}
}
func (h *UserHandler) Login(c *gin.Context) {
	req := dto.LoginRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	tokenString, err := h.userUsecase.Login(req)
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
		Result: tokenString,
	}
	c.JSON(http.StatusOK, data)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	req := dto.CreateUserRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		data := dto.BaseResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, data)
		return
	}
	res, err := h.userUsecase.CreateUser(req)
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

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	req := dto.UpdateProfileRequest{}
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
	file, _ := c.FormFile("file")
	var ioFile multipart.File
	ioFile = nil
	if file != nil {
		ioFile, err = file.Open()
		if err != nil {

			data := dto.BaseResponse{
				Status: http.StatusBadRequest,
				Error:  err.Error(),
			}
			c.JSON(http.StatusBadRequest, data)
			return
		}
	}

	res, err := h.userUsecase.UpdateProfile(req, user, ioFile)
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

func (h *UserHandler) ChangePassWord(c *gin.Context) {
	req := dto.ChangePassWordRequest{}
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

	res, err := h.userUsecase.ChangePassWord(req, user)
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
func (h *UserHandler) GetProfile(c *gin.Context) {

	user := middleware.GetUserFromContext(c)
	res := h.userUsecase.GetProfile(user)
	data := dto.BaseResponse{
		Status: http.StatusOK,
		Result: res,
	}
	c.JSON(http.StatusOK, data)
}
