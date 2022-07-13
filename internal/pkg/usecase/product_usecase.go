package usecase

import (
	"api/internal/pkg/domain/domain_model/dto"
	"api/internal/pkg/domain/domain_model/entity"
	"api/internal/pkg/domain/repository"
	"api/internal/pkg/service"
	"api/pkg/infrastucture/db"
	"api/pkg/share/utils"
	"errors"
	"io"
)

type ProductUsecase interface {
	GetProductList(dto.GetProductListRequest) (dto.GetProductListResponse, error)
	GetProductInfo(productID int) (dto.ProductResponse, error)
	UpdateProduct(dto.UpdateProductRequest, io.Reader) (dto.ProductResponse, error)
	CreateProduct(dto.CreateProductRequest, io.Reader) (dto.ProductResponse, error)
	DeleteProduct(productID int) error
}
type productUsecase struct {
	ProductRepo service.ProductRepository
}

func (u *productUsecase) GetProductList(req dto.GetProductListRequest) (dto.GetProductListResponse, error) {
	const PageSize = 10
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	products, err := u.ProductRepo.FindProductListWithPagination(
		entity.Product{
			BrandID:    req.BrandID,
			CategoryID: req.CategoryID,
			Name:       req.ProductName,
		}, req.PageNum, PageSize)
	if err != nil {
		return dto.GetProductListResponse{}, err
	}
	productsListRes := []dto.ProductResponse{}
	for _, product := range products {
		productsListRes = append(productsListRes, dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Discount:    product.Discount,
			Quantity:    product.Quantity,
			Description: product.Description,
			ImgURL:      product.ImgURL,
			BrandProduct: dto.BrandResponse{
				ID:    product.BrandProduct.ID,
				Title: product.BrandProduct.Title,
			},
			CategoryProduct: dto.CategoryResponse{
				ID:    product.CategoryProduct.ID,
				Title: product.CategoryProduct.Title,
			},
		})
	}
	return dto.GetProductListResponse{
		PageNum:     req.PageNum,
		PageSize:    PageSize,
		ProductList: productsListRes,
	}, nil
}
func (u *productUsecase) GetProductInfo(productID int) (dto.ProductResponse, error) {
	product, err := u.ProductRepo.FirstProduct(entity.Product{ID: productID})
	if err != nil {
		return dto.ProductResponse{}, err
	}
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Discount:    product.Discount,
		Quantity:    product.Quantity,
		Description: product.Description,
		ImgURL:      product.ImgURL,
		BrandProduct: dto.BrandResponse{
			ID:    product.BrandProduct.ID,
			Title: product.BrandProduct.Title,
		},
		CategoryProduct: dto.CategoryResponse{
			ID:    product.CategoryProduct.ID,
			Title: product.CategoryProduct.Title,
		},
	}, nil
}
func (u *productUsecase) UpdateProduct(req dto.UpdateProductRequest, imgProduct io.Reader) (dto.ProductResponse, error) {
	product, err := u.ProductRepo.FirstProduct(entity.Product{ID: req.ID})
	if err != nil {
		return dto.ProductResponse{}, err
	}
	if product.ID == 0 {
		return dto.ProductResponse{}, errors.New("product not exist")
	}
	newProduct := entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Discount:    req.Discount,
		BrandID:     req.BrandID,
		CategoryID:  req.CategoryID,
		Quantity:    req.Quantity,
	}
	if imgProduct != nil {
		url, errUpload := utils.UploadFile(imgProduct, product.Name, []string{"product"})
		if errUpload != nil {
			return dto.ProductResponse{}, errUpload
		}
		newProduct.ImgURL = url
	}
	_, err = u.ProductRepo.UpdateProduct(newProduct, product)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	product, err = u.ProductRepo.FirstProduct(entity.Product{ID: req.ID})
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Discount:    product.Discount,
		Quantity:    product.Quantity,
		Description: product.Description,
		ImgURL:      product.ImgURL,
		BrandProduct: dto.BrandResponse{
			ID:    product.BrandProduct.ID,
			Title: product.BrandProduct.Title,
		},
		CategoryProduct: dto.CategoryResponse{
			ID:    product.CategoryProduct.ID,
			Title: product.CategoryProduct.Title,
		},
	}, err
}

func (u *productUsecase) CreateProduct(req dto.CreateProductRequest, imgProduct io.Reader) (dto.ProductResponse, error) {
	product := entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Discount:    req.Discount,
		BrandID:     req.BrandID,
		CategoryID:  req.CategoryID,
		Quantity:    req.Quantity,
	}
	if imgProduct != nil {
		url, errUpload := utils.UploadFile(imgProduct, product.Name, []string{"product"})
		if errUpload != nil {
			return dto.ProductResponse{}, errUpload
		}
		product.ImgURL = url
	}
	product, err := u.ProductRepo.CreateProduct(product)

	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Discount:    product.Discount,
		Quantity:    product.Quantity,
		Description: product.Description,
		ImgURL:      product.ImgURL,
		BrandProduct: dto.BrandResponse{
			ID:    product.BrandProduct.ID,
			Title: product.BrandProduct.Title,
		},
		CategoryProduct: dto.CategoryResponse{
			ID:    product.CategoryProduct.ID,
			Title: product.CategoryProduct.Title,
		},
	}, err
}
func (u *productUsecase) DeleteProduct(productID int) error {

	product, err := u.ProductRepo.FirstProduct(entity.Product{ID: productID})
	if err != nil {
		return err
	}
	if product.ID == 0 {
		return errors.New("product not exist")
	}
	err = u.ProductRepo.DeleteProduct(product)
	return err
}
func NewProductUsecase(db db.Database) *productUsecase {
	repo := repository.NewProductRepository(db)
	return &productUsecase{
		ProductRepo: repo,
	}
}
