package http

import (
	"encoding/json"
	"html/template"
	"net/http"
)

var jsonerr JSONError

type JSONError struct {
	Succes string `json:"succes"`
	Error  string `json:"error,omitempty"`
}

func (h *handler) ProductsHandle(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("front/product.html")
	if err != nil {
		jsonerr = JSONError{Succes: "false", Error: err.Error()}
		json.NewEncoder(w).Encode(jsonerr)
	}
	err = temp.ExecuteTemplate(w, "product", nil)
	if err != nil {
		jsonerr = JSONError{Succes: "false", Error: err.Error()}
		json.NewEncoder(w).Encode(jsonerr)
	}

	w.WriteHeader(http.StatusOK)
	jsonerr = JSONError{Succes: "true"}
	json.NewEncoder(w).Encode(jsonerr)
}

func (h *handler) Admin(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("front/adminLogin.html")
	if err != nil {
		jsonerr = JSONError{Succes: "false", Error: err.Error()}
		json.NewEncoder(w).Encode(jsonerr)
	}
	err = temp.ExecuteTemplate(w, "adminLogin", nil)
	if err != nil {
		jsonerr = JSONError{Succes: "false", Error: err.Error()}
		json.NewEncoder(w).Encode(jsonerr)
	}

	w.WriteHeader(http.StatusOK)
	jsonerr = JSONError{Succes: "true"}
	json.NewEncoder(w).Encode(jsonerr)
}

func (h *handler) AdminHome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("front/admin.html")
	if err != nil {
		jsonerr = JSONError{Succes: "false", Error: err.Error()}
		json.NewEncoder(w).Encode(jsonerr)
	}
	err = temp.ExecuteTemplate(w, "admin", nil)
	if err != nil {
		jsonerr = JSONError{Succes: "false", Error: err.Error()}
		json.NewEncoder(w).Encode(jsonerr)
	}

	w.WriteHeader(http.StatusOK)
	jsonerr = JSONError{Succes: "true"}
	json.NewEncoder(w).Encode(jsonerr)
}
