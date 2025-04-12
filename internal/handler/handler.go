package handler

import (
	"encoding/json"
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
	if r.Method != http.MethodGet || r.URL.Path != "/" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "internal/view/index.html")
	return
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || r.URL.Path[:3] != "/r/" && len(r.URL.Path[3:]) != 8 {
		http.Error(w, "Adress not exists", http.StatusNotFound)
		return
	}

	//ordinary := ""

	// get original link from storage and redirect
}

func (h *Handler) EncodeURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.URL.Path != "/api/encode" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var u model.URL
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.Unireslocator = h.Service.EncodeAndSaveURL(u.Unireslocator)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}
