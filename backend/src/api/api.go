package api

import (
	"encoding/json"
	"github.com/gsistelos/todo-app/db"
	"golang.org/x/crypto/bcrypt"
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

	router.HandleFunc("POST /users", makeHandlerFunc(s.handleCreateUser))
	router.HandleFunc("GET /users", makeHandlerFunc(s.handleGetUsers))
	router.HandleFunc("GET /users/{id}", makeHandlerFunc(s.handleGetUserByID))
	router.HandleFunc("PUT /users/{id}", makeHandlerFunc(s.handleUpdateUser))
	router.HandleFunc("DELETE /users/{id}", makeHandlerFunc(s.handleDeleteUser))

	http.ListenAndServe(s.listenAddr, router)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}
}
