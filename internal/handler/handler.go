package handler

import (
	"fmt"
	"net/http"
	"ushrt/internal/service"
)

type Handler struct {
	*service.Service
}

func New(s *service.Service) Handler {
	return Handler{
		s,
	}
}

func (h Handler) LoadView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(r.URL.Path) == 9 {
		h.Redirect(w, r)
		return
	}

	http.ServeFile(w, r, "internal/view/index.html")
	return
}

func (h Handler) EncodeURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi from encode")
}

func (h Handler) DecodeUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi from decode")
}

func (h Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi from redirect")
}
