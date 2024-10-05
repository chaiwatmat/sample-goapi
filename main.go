package main

import (
    "fmt"
    "net/http"
    "main/controllers"
)

// setupRoutes initializes the URL routing
func setupRoutes() {
    http.HandleFunc("/hello", controllers.HelloController)
    http.HandleFunc("/user", controllers.GetUserController)
}

func main() {
    setupRoutes()

    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
