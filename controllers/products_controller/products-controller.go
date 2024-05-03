package products_controller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/rofinp/go-web-native/models"
	"github.com/rofinp/go-web-native/services/category_service"
	"github.com/rofinp/go-web-native/services/product_service"
)

func Products(w http.ResponseWriter, r *http.Request) {
	product, err := product_service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"products": product,
	}

	/* Teknik Path Relatif */
	// Parsing template dari file HTML terpisah (contoh: index.html)
	tmpl, err := template.ParseFiles("../../views/product/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Eksekusi template dan tulis ke ResponseWriter
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("../../views/product/create.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		categories, err := category_service.GetAllCategories()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]any{
			"categories": categories,
		}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case "POST":
		id, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Create a new product object with name from form value
		product := models.Product{
			Name: r.FormValue("name"),
			Category: models.Category{
				ID: uint(id),
			},
			Stock:       int(stock),
			Description: r.FormValue("description"),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err = product_service.CreateProduct(product)
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func DetailProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := product_service.DetailProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"product": product,
	}

	tmpl, err := template.ParseFiles("../../views/product/detail.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("../../views/product/edit.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := product_service.DetailProduct(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		categories, err := category_service.GetAllCategories()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]any{
			"categories": categories,
			"product":    product,
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, data)

	case "POST":
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		categoryID, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product := models.Product{
			Name: r.FormValue("name"),
			Category: models.Category{
				ID: uint(categoryID),
			},
			Stock:       int(stock),
			Description: r.FormValue("description"),
			UpdatedAt:   time.Now(),
		}

		if err := product_service.UpdateProduct(id, product); err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := product_service.DeleteProduct(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
