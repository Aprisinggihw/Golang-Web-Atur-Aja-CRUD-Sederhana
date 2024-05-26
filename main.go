package main

import (
	"golang-web-crud/config"
	"golang-web-crud/controllers/categorycontroller"
	"golang-web-crud/controllers/homecontroller"
	"golang-web-crud/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)
	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	//Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
