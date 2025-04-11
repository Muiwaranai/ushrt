package handler

import (
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

	http.ServeFile(w, r, "internal/view/index.html")
	return
}
