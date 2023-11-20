package categorycontroller

import (
	"net/http"
	"quest-rolling-system-RnD-backup/entities"
	"quest-rolling-system-RnD-backup/models/categorymodel"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("page/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("page/category/add.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category
		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		// Jika gagal
		if success := categorymodel.Add(category); !success {
			temp, _ := template.ParseFiles("page/category/add.html")
			temp.Execute(w, nil)
		}

		// Jika berhasil
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("page/category/edit.html")
		if err != nil {
			panic(err)
		}

		// Secara default tipe datanya string
		idString := r.URL.Query().Get("id")
		// Konversi string ke int
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category := categorymodel.Detail(id)
		data := map[string]any{
			"category": category,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var category entities.Category
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		if success := categorymodel.Update(id, category); !success {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	if err := categorymodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		query := r.FormValue("query")
		results := categorymodel.Search(query)
		data := map[string]any{
			"categories": results,
		}
		temp, err := template.ParseFiles("page/category/index.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)
	}
}