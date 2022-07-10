package usecase

import (
	"api/internal/pkg/domain/domain_model/dto"
	"api/internal/pkg/domain/domain_model/entity"
	"api/internal/pkg/domain/repository"
	"api/internal/pkg/service"
	"api/pkg/infrastucture/db"
)

type BrandUsecase interface {
	GetBrandList() (dto.GetBrandListResponse, error)
	UpdateBrand(dto.UpdateBrandRequest) (dto.BrandResponse, error)
	CreateBrand(dto.CreateBrandRequest) (dto.BrandResponse, error)
	DeleteBrand(brandID int) error
}
type brandUsecase struct {
	brandRepo service.BrandRepository
}

func (u *brandUsecase) GetBrandList() (dto.GetBrandListResponse, error) {
	brandList, err := u.brandRepo.FindBrandList(entity.Brand{})
	if err != nil {
		return dto.GetBrandListResponse{}, err
	}
	brandResList := []dto.BrandResponse{}
	for _, brand := range brandList {
		brandResList = append(brandResList, dto.BrandResponse{Title: brand.Title, ID: brand.ID})
	}
	return dto.GetBrandListResponse{
		Total:     len(brandList),
		BrandList: brandResList,
	}, nil
}
func (u *brandUsecase) UpdateBrand(req dto.UpdateBrandRequest) (dto.BrandResponse, error) {
	_, err := u.brandRepo.UpdateBrand(entity.Brand{
		Title: req.Title,
	}, entity.Brand{
		ID: req.ID,
	})
	return dto.BrandResponse{
		ID:    req.ID,
		Title: req.Title,
	}, err
}

func (u *brandUsecase) CreateBrand(req dto.CreateBrandRequest) (dto.BrandResponse, error) {
	brand, err := u.brandRepo.CreateBrand(entity.Brand{
		Title: req.Title,
	})
	return dto.BrandResponse{
		ID:    brand.ID,
		Title: brand.Title,
	}, err
}
func (u *brandUsecase) DeleteBrand(brandID int) error {
	err := u.brandRepo.DeleteBrand(entity.Brand{
		ID: brandID,
	})
	return err
}
func NewBrandUsecase(db db.Database) *brandUsecase {
	repo := repository.NewBrandRepository(db)
	return &brandUsecase{
		brandRepo: repo,
	}
}
