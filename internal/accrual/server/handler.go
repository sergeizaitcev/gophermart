package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sergeizaitcev/gophermart/internal/accrual/service"
)

// handler определяет HTTP-обработчик для accrual
// реализует интерфейс http.Handler
type handler struct {
	mux *chi.Mux
	service *service.Service
}

// NewHandler возвращает новый экземпляр handler
func NewHandler(s *service.Service) http.Handler {
	r := &handler{
		mux: chi.NewRouter(),
		service: s,
	}
	r.init()
	return r
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *handler) init() {
	h.mux.Route("/api", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/orders", h.registerOrder)
			r.Post("/goods", h.createMatch)
			r.Get("/orders/{number}", h.getOrder)
		})
	})
}

func (h *handler) registerOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusTeapot)

	json.NewEncoder(w).Encode(map[string]string{"response": "I'm not implemented yet, but someday it will happen"})
}

func (h *handler) createMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusTeapot)

	json.NewEncoder(w).Encode(map[string]string{"response": "I'm not implemented yet, but someday it will happen"})
}

func (h *handler) getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusTeapot)

	json.NewEncoder(w).Encode(map[string]string{"response": "I'm not implemented yet, but someday it will happen"})
}