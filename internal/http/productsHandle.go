package http

import "net/http"

func (h *handler) ProductsHandle(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "template.html", nil)
}
