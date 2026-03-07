package web

import "time"

type Order struct {
	ID           int         `json:"id"`
	CustomerName string      `json:"customerName"`
	Items        []OrderItem `json:"items"`
	Total        int         `json:"total"`
	Payment      string      `json:"payment"`
	Status       string      `json:"status"`
	CreatedAt    time.Time   `json:"createdAt"`
}

type OrderItem struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}