package repository

import (
	"context"
	"database/sql"
	"golang-api/domain"
	"golang-api/helper"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	FindAll(ctx context.Context, db *sql.DB) []domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, productId int)
}

type productRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &productRepositoryImpl{}
}

func (r *productRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	result, err := tx.ExecContext(ctx,
		"INSERT INTO nicesu(name, price, image) VALUES (?, ?, ?)",
		product.Name, product.Price, product.Image,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (r *productRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {

	_, err := tx.ExecContext(ctx,
		"UPDATE nicesu SET name = ?, price = ?, image = ? WHERE id = ?",
		product.Name, product.Price, product.Image, product.Id,
	)

	if err != nil {
		panic(err)
	}

	return product
}

func (r *productRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId int) {
	SQL := "DELETE FROM nicesu WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, productId)
	helper.PanicIfError(err)
}

func (r *productRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.Product {
	rows, err := db.QueryContext(ctx, "SELECT id, name, price, image FROM nicesu")
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Image)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	return products
}