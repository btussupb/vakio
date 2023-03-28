package http

import (
	"fmt"
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
			fmt.Println("ParseFiles")
		}
		err = temp.ExecuteTemplate(w, "index", nil)
		if err != nil {
			fmt.Println("ExecuteTemplate from method Get")
		}
	case http.MethodPost:
		h.methodPost(w, r)
	}
}

func (h *handler) methodPost(w http.ResponseWriter, r *http.Request) {
	var userInput mngSrv.User
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	// fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	userInput.Name = r.FormValue("name")
	userInput.Number = r.FormValue("number")
	userInput.City = r.FormValue(("sity"))

	// fmt.Println(r.PostForm)

	// if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
	// 	fmt.Println("Decode r.Body:", err)
	// }
	fmt.Println(userInput)
	// defer r.Body.Close()

	fmt.Println(h)
	if err := h.mngSrv.PostUser(userInput); err != nil {
		// w.Write([]byte("ошибка на сервере"))
	} else {
		// w.Write([]byte("мы вам напишем"))
	}
}
