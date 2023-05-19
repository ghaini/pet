package http

import (
	"github.com/go-chi/chi"
)

func Routes(r chi.Router, h *Handler) {
	r.Route("/pets", func(r chi.Router) {
		r.Get("/", h.ListPets)
		r.Post("/", h.AddPet)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetPet)
			r.Delete("/", h.DeletePet)
		})
	})
}
