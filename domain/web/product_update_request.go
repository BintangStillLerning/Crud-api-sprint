package web

type ProductUpdateRequest struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
    Image string 
}