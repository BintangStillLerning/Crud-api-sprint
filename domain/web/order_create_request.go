package web

import "time"

type Order struct {
	ID           int         `json:"id"`
	CustomerName string      `json:"customer_name"`
	Items        []OrderItem `json:"items"`
	Total        int         `json:"total"`
	Payment      string      `json:"payment"`
	Status       string      `json:"status"`
	CreatedAt    time.Time   `json:"created_at"`
}

type OrderItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}