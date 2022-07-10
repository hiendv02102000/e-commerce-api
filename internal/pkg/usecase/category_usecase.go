package usecase

import (
	"api/internal/pkg/domain/domain_model/dto"
	"api/internal/pkg/domain/domain_model/entity"
	"api/internal/pkg/domain/repository"
	"api/internal/pkg/service"
	"api/pkg/infrastucture/db"
)

type CategoryUsecase interface {
	GetCategoryList() (dto.GetCategoryListResponse, error)
	UpdateCategory(dto.UpdateCategoryRequest) (dto.CategoryResponse, error)
	CreateCategory(dto.CreateCategoryRequest) (dto.CategoryResponse, error)
	DeleteCategory(CategoryID int) error
}
type categoryUsecase struct {
	CategoryRepo service.CategoryRepository
}

func (u *categoryUsecase) GetCategoryList() (dto.GetCategoryListResponse, error) {
	CategoryList, err := u.CategoryRepo.FindCategoryList(entity.Category{})
	if err != nil {
		return dto.GetCategoryListResponse{}, err
	}
	CategoryResList := []dto.CategoryResponse{}
	for _, Category := range CategoryList {
		CategoryResList = append(CategoryResList, dto.CategoryResponse{Title: Category.Title, ID: Category.ID})
	}
	return dto.GetCategoryListResponse{
		Total:        len(CategoryList),
		CategoryList: CategoryResList,
	}, nil
}
func (u *categoryUsecase) UpdateCategory(req dto.UpdateCategoryRequest) (dto.CategoryResponse, error) {
	_, err := u.CategoryRepo.UpdateCategory(entity.Category{
		Title: req.Title,
	}, entity.Category{
		ID: req.ID,
	})
	return dto.CategoryResponse{
		ID:    req.ID,
		Title: req.Title,
	}, err
}

func (u *categoryUsecase) CreateCategory(req dto.CreateCategoryRequest) (dto.CategoryResponse, error) {
	Category, err := u.CategoryRepo.CreateCategory(entity.Category{
		Title: req.Title,
	})
	return dto.CategoryResponse{
		ID:    Category.ID,
		Title: Category.Title,
	}, err
}
func (u *categoryUsecase) DeleteCategory(CategoryID int) error {
	err := u.CategoryRepo.DeleteCategory(entity.Category{
		ID: CategoryID,
	})
	return err
}
func NewCategoryUsecase(db db.Database) *categoryUsecase {
	repo := repository.NewCategoryRepository(db)
	return &categoryUsecase{
		CategoryRepo: repo,
	}
}
