package api

import (
	"encoding/json"
	"github.com/gsistelos/todo-app/db"
	"net/http"
)

type APIServer struct {
	listenAddr string
	db         *db.MysqlDB
}

type apiError struct {
	Error string `json:"error"`
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func NewAPIServer(listenAddr string, db *db.MysqlDB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (s *APIServer) Run() {
	router := http.NewServeMux()

	router.HandleFunc("POST /users", defaultHandler(s.handleCreateUser))
	router.HandleFunc("GET /users", defaultHandler(s.handleGetUsers))
	router.HandleFunc("GET /users/{id}", defaultHandler(s.handleGetUserByID))
	router.HandleFunc("PUT /users/{id}", defaultHandler(s.handleUpdateUser))
	router.HandleFunc("DELETE /users/{id}", defaultHandler(s.handleDeleteUser))

	router.HandleFunc("POST /login", defaultHandler(s.handleLogin))

	http.ListenAndServe(s.listenAddr, router)
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func defaultHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}
}
