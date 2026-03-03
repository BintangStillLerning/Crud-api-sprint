package main

import (
	"database/sql"
	"golang-api/controller"
	"golang-api/repository"
	"golang-api/service"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	// ================= DATABASE =================
	dbURL := os.Getenv("root:password@tcp(containers-us-west-xxx.railway.app:6543)/railway")

	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB ERROR:", err)
	}

	validate := validator.New()

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := httprouter.New()

	router.POST("/products", productController.Create)
	router.GET("/products", productController.FindAll)
	router.DELETE("/products/:productId", productController.Delete)
	router.PUT("/products/:productId", productController.Update)

	// ================= PORT =================
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}