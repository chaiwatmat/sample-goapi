package controllers

import (
    "encoding/json"
    "net/http"
    "main/models"
)

// GetUserController handles GET requests to fetch user information
func GetUserController(w http.ResponseWriter, r *http.Request) {
    user := models.User{ID: 1, Name: "Alice", Age: 30}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

