package categories_controller

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/rofinp/go-web-native/models"
	"github.com/rofinp/go-web-native/services/category_service"
)

func Categories(w http.ResponseWriter, r *http.Request) {
	category, err := category_service.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"categories": category,
	}

	/* Teknik Path Relatif */
	// Parsing template dari file HTML terpisah (contoh: index.html)
	tmpl, err := template.ParseFiles("../../views/category/index.html")
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

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Parse the "create.html" template file
		tmpl, err := template.ParseFiles("../../views/category/create.html")
		if err != nil {
			// Return internal server error if template parsing fails
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set response header and write 200 OK status code
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		// Execute the template with no data
		tmpl.Execute(w, nil)

	case "POST":
		// Create a new Category object with name from form value
		category := models.Category{
			Name:      r.FormValue("name"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// Create the category in the database
		err := category_service.CreateCategory(category)
		if err != nil {
			// Return bad request status code with error message if creation fails
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Redirect to "/categories" page with 303 See Other status code
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func EditCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Parse the edit.html template file
		tmpl, err := template.ParseFiles("../../views/category/edit.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Retrieve the category ID from the URL query parameters
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Fetch the category details from the database
		category, err := category_service.DetailsCategory(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create a data map with the category details
		data := map[string]any{
			"category": category,
		}

		// Render the template with the data
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, data)

	case "POST":
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Create a new Category object with the updated name and current timestamp
		category := models.Category{
			Name:      r.FormValue("name"),
			UpdatedAt: time.Now(),
		}

		// Update the category details in the database
		if err := category_service.UpdateCategory(id, category); err != nil {
			// Redirect the user to the previous page if there is an error
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		// Redirect the user to the category list page
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// Retrieve the category ID from the URL query parameters and
	// Convert the ID string to an integer and check for any errors
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete the category with the given ID
	if err := category_service.DeleteCategory(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect the user to the "/categories" page
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
