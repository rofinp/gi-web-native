package home_controller

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	/* Teknik Path Absolute */
	// Mendapatkan direktori kerja saat ini
	currentDir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Menggabungkan direktori kerja dengan path file HTML
	htmlFilePath := filepath.Join(currentDir, "../../views/home/index.html")

	// Parsing template dari file HTML terpisah (contoh: index.html)
	tmpl, err := template.ParseFiles(htmlFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Eksekusi template dan tulis ke ResponseWriter
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
