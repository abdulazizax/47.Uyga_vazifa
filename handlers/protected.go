package handlers

import (
	"fmt"
	"net/http"
)

type key int

const (
	roleKey key = iota
)

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(roleKey).(string)
	if role != "admin" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "Welcome, admin!")
}
