package http

import "net/http"

func NewHandler() handler {
	var h handler
	return h
}

func (h *handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.RootHandle)
	mux.HandleFunc("/products", h.ProductsHandle)
	mux.Handle("/front/css/", http.StripPrefix("/front/css/", http.FileServer(http.Dir("./front/css"))))
	return mux
}
