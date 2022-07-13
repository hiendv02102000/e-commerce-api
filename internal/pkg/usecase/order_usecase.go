package usecase

import (
	"api/internal/pkg/domain/domain_model/dto"
	"api/internal/pkg/domain/domain_model/entity"
	"api/internal/pkg/domain/repository"
	"api/internal/pkg/service"
	"api/pkg/infrastucture/db"
	"api/pkg/share/middleware"
	"errors"
	"fmt"
)

type OrderUsecase interface {
	GetOrderList(dto.GetOrderListRequest) (dto.GetOrderListResponse, error)
	GetOrderInfo(transactionCode string, idCustomer int) (dto.OrderResponse, error)
	ChangeStatusOrder(dto.UpdateStatucOrderRequest) (dto.OrderResponse, error)
	// customer
	GetCart(user entity.Users) (dto.OrderResponse, error)
	CreateCart(user entity.Users) (dto.OrderResponse, error) //
	PutProductToCart(req dto.PutProductToCartRequest) error
	DeleteProductFromCart(dto.DeleteProductFromCartRequest) error //
	CancelOrder(transactionCode string, idCustomer int) error
	ConfirmOrder(user entity.Users) error
}
type orderUsecase struct {
	OrderRepo           service.OrderRepository
	ProductRepo         service.ProductRepository
	SelectedProductRepo service.SelectedProductRepository
}

func (u *orderUsecase) GetOrderList(req dto.GetOrderListRequest) (dto.GetOrderListResponse, error) {
	const PageSize = 10
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	orders, err := u.OrderRepo.FindOrderListWithPagination(entity.Order{
		TransactionCode: req.TransactionCode,
		Address:         req.Address,
		Phone:           req.Phone,
		Status:          entity.OrderStatus(req.OrderStatus),
	}, req.PageNum, PageSize)
	if err != nil {
		return dto.GetOrderListResponse{}, err
	}
	orderListRes := []dto.OrderResponse{}
	for _, order := range orders {
		orderListRes = append(orderListRes, dto.OrderResponse{
			ID:              order.ID,
			TransactionCode: order.TransactionCode,
			TotalPrice:      order.TotalPrice,
			OrderStatus:     order.Address,
			Address:         order.Address,
			Phone:           order.Phone,
			CustomerID:      order.CustomerID,
		})
	}
	return dto.GetOrderListResponse{
		PageNum:   req.PageNum,
		PageSize:  PageSize,
		OrderList: orderListRes,
	}, nil
}
func (u *orderUsecase) GetOrderInfo(transactionCode string, idCustomer int) (dto.OrderResponse, error) {
	order, err := u.OrderRepo.FirstOrder(entity.Order{
		TransactionCode: transactionCode,
	})
	if err != nil {
		return dto.OrderResponse{}, err
	}
	if idCustomer != 0 {
		if idCustomer != order.CustomerID {
			return dto.OrderResponse{}, errors.New("method not allowed")
		}
	}
	productListRes := []dto.ProductResponse{}
	for _, product := range order.SelectedProductList {
		product.ProductInfo, err = u.ProductRepo.FirstProduct(entity.Product{ID: product.ProductID})
		if err != nil {
			return dto.OrderResponse{}, err
		}
		productListRes = append(productListRes, dto.ProductResponse{
			ID:          product.ProductID,
			Name:        product.ProductInfo.Name,
			Price:       product.Price,
			Discount:    product.ProductInfo.Discount,
			Quantity:    product.Quantity,
			Description: product.ProductInfo.Description,
			ImgURL:      product.ProductInfo.ImgURL,
			BrandProduct: dto.BrandResponse{
				ID:    product.ProductInfo.BrandProduct.ID,
				Title: product.ProductInfo.BrandProduct.Title,
			},
			CategoryProduct: dto.CategoryResponse{
				ID:    product.ProductInfo.CategoryProduct.ID,
				Title: product.ProductInfo.CategoryProduct.Title,
			},
		})
	}
	return dto.OrderResponse{
		ID:              order.ID,
		TransactionCode: order.TransactionCode,
		TotalPrice:      order.TotalPrice,
		OrderStatus:     order.Address,
		Address:         order.Address,
		Phone:           order.Phone,
		CustomerID:      order.CustomerID,
		ProductList:     productListRes,
	}, nil
}
func (u *orderUsecase) ChangeStatusOrder(req dto.UpdateStatucOrderRequest) (dto.OrderResponse, error) {
	order, err := u.OrderRepo.FirstOrder(entity.Order{
		TransactionCode: req.TransactionCode,
	})
	if err != nil {
		return dto.OrderResponse{}, err
	}
	if order.ID == 0 {
		return dto.OrderResponse{}, errors.New("order not exist")
	}
	if order.Status == entity.CART || order.Status == entity.CANCLED {
		return dto.OrderResponse{}, errors.New("method not allowed")
	}
	newOrder := order
	newOrder.Status = entity.OrderStatus(req.OrderStatus)
	order, err = u.OrderRepo.UpdateOrder(newOrder, order)
	if err != nil {
		return dto.OrderResponse{}, err
	}
	return dto.OrderResponse{
		ID:              order.ID,
		TransactionCode: order.TransactionCode,
		TotalPrice:      order.TotalPrice,
		OrderStatus:     order.Address,
		Address:         order.Address,
		Phone:           order.Phone,
		CustomerID:      order.CustomerID,
	}, nil
}
func (u *orderUsecase) CreateCart(user entity.Users) (dto.OrderResponse, error) {
	order, err := u.OrderRepo.FirstOrder(entity.Order{
		CustomerID: user.ID,
		Status:     entity.CART,
	})
	if err != nil {
		return dto.OrderResponse{}, err
	}
	productListRes := []dto.ProductResponse{}
	if order.ID == 0 {
		code, err := middleware.GenerateTransactionCode()
		if err != nil {
			return dto.OrderResponse{}, err
		}
		order, err = u.OrderRepo.CreateOrder(entity.Order{
			TransactionCode: code,
			CustomerID:      user.ID,
			Status:          entity.CART,
		})
		if err != nil {
			return dto.OrderResponse{}, err
		}
	} else {
		return dto.OrderResponse{}, errors.New("cart already exist")
	}
	return dto.OrderResponse{
		ID:              order.ID,
		TransactionCode: order.TransactionCode,
		TotalPrice:      order.TotalPrice,
		OrderStatus:     order.Address,
		Address:         order.Address,
		Phone:           order.Phone,
		CustomerID:      order.CustomerID,
		ProductList:     productListRes,
	}, nil
}
func (u *orderUsecase) GetCart(user entity.Users) (dto.OrderResponse, error) {
	order, err := u.OrderRepo.FirstOrder(entity.Order{
		CustomerID: user.ID,
		Status:     entity.CART,
	})
	if err != nil {
		return dto.OrderResponse{}, err
	}
	productListRes := []dto.ProductResponse{}
	if order.ID == 0 {
		return dto.OrderResponse{}, errors.New("cart not exist")
	} else {
		for _, product := range order.SelectedProductList {
			product.ProductInfo, err = u.ProductRepo.FirstProduct(entity.Product{ID: product.ProductID})
			if err != nil {
				return dto.OrderResponse{}, err
			}
			productListRes = append(productListRes, dto.ProductResponse{
				ID:          product.ProductID,
				Name:        product.ProductInfo.Name,
				Price:       product.Price,
				Discount:    product.ProductInfo.Discount,
				Quantity:    product.Quantity,
				Description: product.ProductInfo.Description,
				ImgURL:      product.ProductInfo.ImgURL,
				BrandProduct: dto.BrandResponse{
					ID:    product.ProductInfo.BrandProduct.ID,
					Title: product.ProductInfo.BrandProduct.Title,
				},
				CategoryProduct: dto.CategoryResponse{
					ID:    product.ProductInfo.CategoryProduct.ID,
					Title: product.ProductInfo.CategoryProduct.Title,
				},
			})
		}
	}
	return dto.OrderResponse{
		ID:              order.ID,
		TransactionCode: order.TransactionCode,
		TotalPrice:      order.TotalPrice,
		OrderStatus:     order.Address,
		Address:         order.Address,
		Phone:           order.Phone,
		CustomerID:      order.CustomerID,
		ProductList:     productListRes,
	}, nil
}
func (u *orderUsecase) PutProductToCart(req dto.PutProductToCartRequest) error {
	order, err := u.OrderRepo.FirstOrder(entity.Order{
		TransactionCode: req.TransactionCode,
		Status:          entity.CART,
	})
	if err != nil {
		return err
	}
	if order.ID == 0 {
		return errors.New("cart not exist")
	}
	productInfo, err := u.ProductRepo.FirstProduct(entity.Product{
		ID: req.ProductID,
	})
	if err != nil {
		return err
	}
	if productInfo.ID == 0 {
		return errors.New("product not exist")
	}
	if req.Quantity > productInfo.Quantity {
		return errors.New("order quantity greater storage quantity " + fmt.Sprintf("(%d>%d)", req.Quantity, productInfo.Quantity))
	}
	for _, selectedProduct := range order.SelectedProductList {
		if selectedProduct.ProductID == req.ProductID {
			_, err = u.SelectedProductRepo.UpdateSelectedProduct(entity.SelectedProduct{
				Price:    productInfo.Price,
				Quantity: req.Quantity,
			}, selectedProduct)
			return err
		}
	}
	_, err = u.SelectedProductRepo.CreateSelectedProduct(entity.SelectedProduct{
		Price:     productInfo.Price,
		Quantity:  req.Quantity,
		OrderID:   order.ID,
		ProductID: productInfo.ID,
	})

	return err
}
func (u *orderUsecase) CancelOrder(transactionCode string, idCustomer int) error {
	order, err := u.OrderRepo.FirstOrder(entity.Order{
		TransactionCode: transactionCode,
	})
	if err != nil {
		return err
	}
	if idCustomer != 0 {
		if idCustomer != order.CustomerID {
			return errors.New("method not allowed")
		}
	}
	_, err = u.OrderRepo.UpdateOrder(entity.Order{Status: entity.CANCLED}, order)
	return err
}
func (u *orderUsecase) ConfirmOrder(user entity.Users) error {
	order, err := u.OrderRepo.FirstOrder(entity.Order{
		CustomerID: user.ID,
		Status:     entity.CART,
	})
	if err != nil {
		return err
	}
	if order.ID == 0 {
		return errors.New("order not exist")
	}
	for _, product := range order.SelectedProductList {
		product.ProductInfo, err = u.ProductRepo.FirstProduct(entity.Product{ID: product.ProductID})
		if err != nil {
			return err
		}
		if product.Quantity > product.ProductInfo.Quantity {
			return errors.New(product.ProductInfo.Name + " order quantity greater storage quantity " + fmt.Sprintf("(%d>%d)", product.Quantity, product.ProductInfo.Quantity))
		}
		product, err = u.SelectedProductRepo.UpdateSelectedProduct(product, entity.SelectedProduct{Price: product.ProductInfo.Price})
		if err != nil {
			return err
		}
	}
	order.CalculatePrice()
	_, err = u.OrderRepo.UpdateOrder(entity.Order{Status: entity.ORDERED}, order)
	return err
}
func (u *orderUsecase) DeleteProductFromCart(req dto.DeleteProductFromCartRequest) error {
	order, err := u.OrderRepo.FirstOrder(entity.Order{
		CustomerID: req.CustomerID,
		Status:     entity.CART,
	})
	if err != nil {
		return err
	}

	if order.ID == 0 {
		return errors.New("cart not exist")
	}
	err = u.SelectedProductRepo.DeleteSelectedProduct(entity.SelectedProduct{
		OrderID:   order.ID,
		ProductID: req.ProductID,
	})
	return err
}
func NewOrderUsecase(db db.Database) *orderUsecase {
	repo := repository.NewOrderRepository(db)
	repoPr := repository.NewProductRepository(db)
	repoPrS := repository.NewSelectedProductRepository(db)
	return &orderUsecase{
		OrderRepo:           repo,
		ProductRepo:         repoPr,
		SelectedProductRepo: repoPrS,
	}
}
