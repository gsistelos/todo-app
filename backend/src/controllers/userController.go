package controllers

import (
	"net/http"
)

func getAllUsers(w http.ResponseWriter, router *http.Request) {
	html := "<html><body><h1>All Users</h1><ul>"

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(html))
}

func getUser(w http.ResponseWriter, router *http.Request) {
	html := "<html><body><h1>User</h1><ul>"

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(html))
}
