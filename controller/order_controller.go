package controller

import (
	"encoding/json"
	"golang-api/domain/web"
	"golang-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(service service.OrderService) OrderController {
	return OrderController{
		orderService: service,
	}
}

func (c OrderController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request web.Order

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := c.orderService.Create(r.Context(), request)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}