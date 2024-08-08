package api

import (
	"net/http"

	"github.com/carlos-hfc/go-react/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
)

type ApiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := ApiHandler{
		q: q,
	}

	r := chi.NewRouter()

	a.r = r
	return a
}
