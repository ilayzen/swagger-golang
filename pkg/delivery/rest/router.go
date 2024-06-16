package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h *Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/static/*", h.GetSwaggerByDir())
	r.Get("/dir/*", h.GetSwaggerByReadFiles())

	return r
}
