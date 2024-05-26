package productcontroller

import (
	"fmt"
	"golang-web-crud/entities"
	"golang-web-crud/models/categorymodel"
	"golang-web-crud/models/productmodel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Error converting id to string func Detail")
		panic(err)
	}
	product := productmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {
		fmt.Println("Couldn't parse template Detail.html")
		panic(err)
	}
	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/add.html")
		if err != nil {
			fmt.Println("error parsing templates from Add")
			panic(err)
		}
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}
		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var product entities.Product

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			fmt.Println("Invalid parsing categoryId")
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			fmt.Println("Invalid parsing stock")
			panic(err)
		}
		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Stock = int64(stock)
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := productmodel.Add(product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			fmt.Println("error parsing templates from Edit")
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println("error converting id to int")
			panic(err)
		}
		product := productmodel.Detail(id)
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"product": product,
		}
		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var product entities.Product
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println("error converting id to int")
			panic(err)
		}


		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			fmt.Println("Invalid parsing categoryId")
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			fmt.Println("Invalid parsing stock")
			panic(err)
		}
		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Stock = int64(stock)
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()

		if ok := productmodel.Update(id, product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println("Error parsing id")
		panic(err)
	}
	if err := productmodel.Delete(id); err != nil {
		fmt.Println("Error deleting product")
		panic(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
