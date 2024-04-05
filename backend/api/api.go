package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gsistelos/todo-app/db"
)

var (
	corsOrigin = "http://localhost:3000"
)

type APIServer struct {
	listenAddr string
	db         *db.MysqlDB
}

type apiError struct {
	Error string `json:"error,omitempty"`
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

	router.HandleFunc("OPTIONS /api/users", corsHandler)
	router.HandleFunc("POST /api/users", defaultHandler(s.handleCreateUser))
	router.HandleFunc("GET /api/users", defaultHandler(s.handleGetUsers))

	router.HandleFunc("OPTIONS /api/users/{userID}", corsAuthHandler)
	router.HandleFunc("GET /api/users/{userID}", s.jwtHandler(s.handleGetUserByID))
	router.HandleFunc("PUT /api/users/{userID}", s.jwtHandler(s.handleUpdateUser))
	router.HandleFunc("DELETE /api/users/{userID}", s.jwtHandler(s.handleDeleteUser))

	router.HandleFunc("OPTIONS /api/login", corsHandler)
	router.HandleFunc("POST /api/login", defaultHandler(s.handleLogin))

	router.HandleFunc("OPTIONS /api/users/{userID}/tasks", corsAuthHandler)
	router.HandleFunc("POST /api/users/{userID}/tasks", s.jwtHandler(s.handleCreateTask))
	router.HandleFunc("GET /api/users/{userID}/tasks", s.jwtHandler(s.handleGetTasks))

	router.HandleFunc("OPTIONS /api/users/{userID}/tasks/{taskID}", corsAuthHandler)
	router.HandleFunc("GET /api/users/{userID}/tasks/{taskID}", s.jwtHandler(s.handleGetTaskByID))
	router.HandleFunc("PUT /api/users/{userID}/tasks/{taskID}", s.jwtHandler(s.handleUpdateTask))
	router.HandleFunc("DELETE /api/users/{userID}/tasks/{taskID}", s.jwtHandler(s.handleDeleteTask))

	log.Fatal(http.ListenAndServe(s.listenAddr, router))
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func defaultHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", corsOrigin)
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
		}
	}
}

func corsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", corsOrigin)
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
}

func corsAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", corsOrigin)
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
}
