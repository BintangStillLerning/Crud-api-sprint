package web

import "time"

type OrderResponse struct {
	ID           int                  `json:"id"`
	CustomerName string               `json:"customerName"`
	Items        []OrderItemResponse  `json:"items"`
	Total        int                  `json:"total"`
	Payment      string               `json:"payment"`
	Status       string               `json:"status"`
	CreatedAt    time.Time            `json:"createdAt"`
}

type OrderItemResponse struct {
	ProductID   int    `json:"productId"`
	ProductName string `json:"productName"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	Subtotal    int    `json:"subtotal"`
}

