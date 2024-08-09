package api

import (
	"net/http"

	"github.com/guilherme-deschamps/React-Go-AMA-Live-Chat/internal/store/pgstore"

	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	q *pgstore.Queries // Poderia ter colocado aqui uma interface ao invés do tipo concreto, para facilitar mudanças de BD
	r *chi.Mux         // Mux = multiplexer
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

// http.Handler é um tipo que consegue responder a HTTP requests
func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}

	r := chi.NewRouter()

	a.r = r
	return a
}
