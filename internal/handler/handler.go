package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

func (s *Handler) LoadView(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || r.URL.Path != "/" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "internal/view/index_copy.html")
	return
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || r.URL.Path[:3] != "/r/" && len(r.URL.Path) != 11 {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url, err := h.Service.Hash2URL(r.URL.Path)
	if err != nil {
		http.Error(w, fmt.Sprintf("500 server error: %s", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "https://"+url, http.StatusFound)
}

func (h *Handler) ProcessUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.URL.Path != "/api/encode" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var url model.URL
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		log.Println("error during json decoding")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	short, err := h.Service.URL2Hash(url.URL)
	if err != nil {
		http.Error(w, fmt.Sprintf("500 server error: %s", err), http.StatusInternalServerError)
		return
	}

	issuer := os.Getenv("ISSUER")
	if issuer == "" {
		log.Fatal("ISSUER environment variable is missing")
	}
	short = fmt.Sprintf("%s%s", issuer, short)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(short)

}
