package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gsistelos/todo-app/db"
	"net/http"
)

type APIServer struct {
	listenAddr string
	db         *db.MysqlDB
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func NewAPIServer(listenAddr string, db *db.MysqlDB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/users", makeHandlerFunc(s.handleCreateUser)).Methods("POST")
	router.HandleFunc("/users", makeHandlerFunc(s.handleGetUsers)).Methods("GET")
	router.HandleFunc("/users/{id}", makeHandlerFunc(s.handleGetUser)).Methods("GET")
	router.HandleFunc("/users/{id}", makeHandlerFunc(s.handleUpdateUser)).Methods("PUT")
	router.HandleFunc("/users/{id}", makeHandlerFunc(s.handleDeleteUser)).Methods("DELETE")

	http.ListenAndServe(s.listenAddr, router)
}

func errJSON(err string) map[string]string {
	return map[string]string{"error": err}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, errJSON(err.Error()))
		}
	}
}
