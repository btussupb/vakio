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

func (h *handler) Admin(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("front/adminLogin.html")
	if err != nil {
		//
	}
	err = temp.ExecuteTemplate(w, "adminLogin", nil)
	if err != nil {
		//
	}
}

func (h *handler) AdminHome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("front/admin.html")
	if err != nil {
		//
	}
	err = temp.ExecuteTemplate(w, "admin", nil)
	if err != nil {
		//
	}
}
