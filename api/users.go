package handler

import (
	"encoding/json"
	"github/gabrielyea/simple-go-api/services"
	"net/http"
)

// Handler is the exported function that Vercel will invoke
func Handler(w http.ResponseWriter, r *http.Request) {
	user := services.GetUser() // Assuming GetUser() is accessible and correctly returns a user object
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
