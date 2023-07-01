package router

import (
	"net/http"

	"solver/internal/handler"

	"github.com/gorilla/mux"
)

func New(h *handler.Handler) http.Handler {
	r := mux.NewRouter()
	r.Use(SetJSONHeader)

	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(SetJSONHeader)

	apiRouter.HandleFunc("/users", h.CreateUserHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/users/money", h.IncreaseBalanceHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc("/solve", h.SolveTheTaskHandler).Methods(http.MethodPost)

	return r
}

func SetJSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
