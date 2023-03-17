package http

import (
	"html/template"
	"net/http"
)

func (h *handler) ProductsHandle(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("front/product.html")
	if err != nil {
		//
	}
	err = temp.ExecuteTemplate(w, "product", nil)
	if err != nil {
		//
	}
}
