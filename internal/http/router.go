package http

import (
	"net/http"

	mngSrv "github.com/btussupb/vakio/internal/manager"
)

type service struct {
	mngSrv mngSrv.Service
}

func NewMngService(mngSrv mngSrv.Service) *service {
	return &service{mngSrv: mngSrv}
}

func (h *handler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/front/", http.StripPrefix("/front/", http.FileServer(http.Dir("./front/"))))

	mux.HandleFunc("/", h.RootHandle)
	mux.HandleFunc("/products", h.ProductsHandle)
	mux.HandleFunc("/about_us", h.AdminHome)
	mux.Handle("/front/css/", http.StripPrefix("/front/css/", http.FileServer(http.Dir("./front/css"))))
	return mux
}
