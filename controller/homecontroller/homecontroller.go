package homecontroller

import (
	"net/http"
	"text/template"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("page/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
