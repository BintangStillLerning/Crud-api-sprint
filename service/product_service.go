package service

import (
	"context"
	"database/sql"
	"golang-api/domain"
	"golang-api/domain/web"
	"golang-api/helper"
	"golang-api/repository"

	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
	Delete(ctx context.Context, productId int)
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
}

type productServiceImpl struct {
	Repository repository.ProductRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewProductService(repo repository.ProductRepository, db *sql.DB, validate *validator.Validate) ProductService {
	return &productServiceImpl{
		Repository: repo,
		DB:         db,
		Validate:   validate,
	}
}

func (s *productServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {

	s.Validate.Struct(request)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer tx.Rollback()

	product := domain.Product{
		Name:  request.Name,
		Price: request.Price,
		Image: request.Image,
	}

	product = s.Repository.Save(ctx, tx, product)

	err = tx.Commit()
	helper.PanicIfError(err)

	return web.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
		Image: product.Image,
	}
}

func (s *productServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer tx.Rollback()

	s.Repository.Delete(ctx, tx, productId)

	err = tx.Commit()
	helper.PanicIfError(err)
}

func (s *productServiceImpl) Update (ctx context.Context, request web.ProductUpdateRequest)web.ProductResponse{
	s.Validate.Struct(request)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer tx.Rollback()

	product := domain.Product{
		Id: request.Id,
		Name: request.Name,
		Price: request.Price,
		Image: request.Image,
	}
	
	product = s.Repository.Update(ctx, tx, product)

	err = tx.Commit()
	helper.PanicIfError(err)
	
	return web.ProductResponse{
		Id: product.Id,
		Name: product.Name,
		Price: product.Price,
		Image: product.Image,
	}
}

func (s *productServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	products := s.Repository.FindAll(ctx, s.DB)

	var responses []web.ProductResponse
	for _, product := range products {
		responses = append(responses, web.ProductResponse{
			Id:    product.Id,
			Name:  product.Name,
			Price: product.Price,
			Image: product.Image,
		})
	}
	return responses
}