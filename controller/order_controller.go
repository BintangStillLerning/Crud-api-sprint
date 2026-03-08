package controller

import (
	"encoding/json"
	"fmt"
	"io"
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
	defer r.Body.Close()

	var request web.Order

	// DEBUG: baca raw body
	body, _ := io.ReadAll(r.Body)
	fmt.Println("RAW BODY:", string(body))

	// decode lagi dari body
	err := json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// DEBUG: cek hasil decode
	fmt.Printf("DECODE RESULT: %+v\n", request)
	fmt.Println("CustomerName:", request.CustomerName)
	fmt.Println("Items:", request.Items)
	fmt.Println("Payment:", request.Payment)

	response := c.orderService.Create(r.Context(), request)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c OrderController) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	response := c.orderService.FindAll(r.Context())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}