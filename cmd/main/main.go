package main

import (
	"log"
	"net/http"

	"github.com/rofinp/go-web-native/configs"
	"github.com/rofinp/go-web-native/controllers/categories_controller"
	"github.com/rofinp/go-web-native/controllers/home_controller"
	"github.com/rofinp/go-web-native/controllers/products_controller"
)

func main() {
	configs.ConnectToDatabase()

	/* Home page */
	http.HandleFunc("/", home_controller.Home)

	/* Categories Page */
	http.HandleFunc("/categories", categories_controller.Categories)
	http.HandleFunc("/categories/create", categories_controller.CreateCategory)
	http.HandleFunc("/categories/edit", categories_controller.EditCategory)
	http.HandleFunc("/categories/delete", categories_controller.DeleteCategory)

	/* Products Page */
	http.HandleFunc("/products", products_controller.Products)
	http.HandleFunc("/products/create", products_controller.CreateProduct)
	http.HandleFunc("/products/detail", products_controller.DetailProduct)
	http.HandleFunc("/products/edit", products_controller.EditProduct)
	http.HandleFunc("/products/delete", products_controller.DeleteProduct)

	log.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
