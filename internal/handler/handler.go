package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ushrt/internal/model"
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

func (h *Handler) LoadView(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/view/index.html")
	return
}

func (h *Handler) EncodeURL(w http.ResponseWriter, r *http.Request) {
	var u model.URL
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {

}
