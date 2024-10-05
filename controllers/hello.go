package controllers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "main/models"
)

// HelloController handles incoming POST requests for greetings
func HelloController(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var msg models.Message
    if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response := fmt.Sprintf("Hello, %s! You said: %s", msg.Name, msg.Text)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"response": response})
}
