package controller

import (
	"encoding/json"
	"golang-api/domain/web"
	"golang-api/helper"
	"golang-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductController struct {
	Service service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{Service: service}
}


func (c *ProductController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var request web.ProductCreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicIfError(err)

	response := c.Service.Create(r.Context(), request)

	webResponse := web.WebReponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *ProductController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	productId := ps.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	var request web.ProductUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	helper.PanicIfError(err)

	request.Id = id // penting

	response := c.Service.Update(r.Context(), request)

	webResponse := web.WebReponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	helper.WriteToResponseBody(w, webResponse)
}


func (c *ProductController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	productId := ps.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	c.Service.Delete(r.Context(), id)

	webResponse := web.WebReponse{
		Code:   200,
		Status: "OK",
		Data:   "Product deleted successfully",
	}

	helper.WriteToResponseBody(w, webResponse)
}


func (c *ProductController) FindAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	responses := c.Service.FindAll(r.Context())

	webResponse := web.WebReponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	}

	helper.WriteToResponseBody(w, webResponse)
}