package cartcontroller

import (
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("page/cart/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
