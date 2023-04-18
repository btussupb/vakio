package http

import (
	"encoding/json"
	"html/template"
	"net/http"

	mngSrv "github.com/btussupb/vakio/internal/manager"
)

type handler struct {
	mngSrv mngSrv.Service
}

func NewHandler(mngSrv mngSrv.Service) *handler {
	return &handler{
		mngSrv: mngSrv,
	}
}

var temp *template.Template

func (h *handler) RootHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		temp, err := template.ParseFiles("front/index.html")
		if err != nil {
			jsonerr = JSONError{Succes: "false", Error: err.Error()}
			json.NewEncoder(w).Encode(jsonerr)
		}
		err = temp.ExecuteTemplate(w, "index", nil)
		if err != nil {
			jsonerr = JSONError{Succes: "false", Error: err.Error()}
			json.NewEncoder(w).Encode(jsonerr)
		}
	case http.MethodPost:
		h.methodPost(w, r)
	}
}

func (h *handler) methodPost(w http.ResponseWriter, r *http.Request) {
	var userInput mngSrv.User
	if err := r.ParseForm(); err != nil {
		jsonerr = JSONError{Succes: "false", Error: err.Error()}
		json.NewEncoder(w).Encode(jsonerr)
	}

	userInput.Name = r.FormValue("name")
	userInput.Number = r.FormValue("number")
	userInput.City = r.FormValue(("sity"))

	if err := h.mngSrv.PostUser(userInput); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonerr = JSONError{Succes: "false", Error: err.Error()}
		json.NewEncoder(w).Encode(jsonerr)
	} else {
		w.WriteHeader(http.StatusOK)
		jsonerr = JSONError{Succes: "true"}
		json.NewEncoder(w).Encode(jsonerr)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
