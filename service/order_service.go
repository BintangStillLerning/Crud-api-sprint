package service

import (
	"context"
	"database/sql"
	"golang-api/domain/web"
	"golang-api/repository"
)

type OrderService interface {
	Create(ctx context.Context, request web.Order) web.OrderResponse
	FindAll(ctx context.Context) []web.OrderResponse
}

type orderServiceImpl struct {
	OrderRepository   repository.OrderRepository
	ProductRepository repository.ProductRepository
	DB                *sql.DB
}

func NewOrderService(
	orderRepo repository.OrderRepository,
	productRepo repository.ProductRepository,
	db *sql.DB,
) OrderService {
	return &orderServiceImpl{
		OrderRepository:   orderRepo,
		ProductRepository: productRepo,
		DB:                db,
	}
}

func (s *orderServiceImpl) Create(ctx context.Context, request web.Order) web.OrderResponse {

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()


	order := web.Order{
		CustomerName: request.CustomerName,
		Payment:      request.Payment,
		Status:       request.Status,
	}

	// Hitung total


	order.Total = request.Total

	// Save order
	order = s.OrderRepository.Save(ctx, tx, order)

	// Save order items
	for _, item := range request.Items {

		product, err := s.ProductRepository.FindById(ctx, tx, item.ProductID)
		if err != nil {
			panic(err)
		}

		orderItem := web.OrderItem{
			ProductID: product.Id,
			Quantity:  item.Quantity,
		}

		s.OrderRepository.SaveItem(ctx, tx, orderItem)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	return web.OrderResponse{
		CustomerName: order.CustomerName,
		Total:        order.Total,
		Payment:      order.Payment,
		Status:       order.Status,
	}
	
}

func (s *orderServiceImpl) FindAll(ctx context.Context) []web.OrderResponse {

	orders := s.OrderRepository.FindAll(ctx, s.DB)

	var responses []web.OrderResponse

	for _, order := range orders {

		responses = append(responses, web.OrderResponse{
			ID:           order.ID,
			CustomerName: order.CustomerName,
			Payment:      order.Payment,
		})
	}

	return responses
}