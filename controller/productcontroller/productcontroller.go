package productcontroller

import (
	"net/http"
	"quest-rolling-system-RnD-backup/entities"
	"quest-rolling-system-RnD-backup/models/categorymodel"
	"quest-rolling-system-RnD-backup/models/productmodel"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("page/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	product := productmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("page/product/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("page/product/add.html")
		if err != nil {
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
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Stock = stock
		product.Description = r.FormValue("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if success := productmodel.Add(product); !success {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("page/product/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		product := productmodel.Detail(id)
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"product":    product,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Stock = stock
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()

		if success := productmodel.Update(id, product); !success {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.FormValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	if err := productmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		query := r.FormValue("query")
		results := productmodel.Search(query)
		data := map[string]any{
			"products": results,
		}
		temp, err := template.ParseFiles("page/product/index.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)
	}
}

