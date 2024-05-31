package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "h47/handlers"
    "h47/middleware"
    "h47/utils"
)

type LoginPayload struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var payload LoginPayload
    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if payload.Username == "user1" && payload.Password == "password" {
        token, err := utils.GenerateJWT(payload.Username, "admin")
        if err != nil {
            http.Error(w, "Could not generate token", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"token": token})
    } else {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    }
}

func main() {
    http.HandleFunc("/login", LoginHandler)
    http.Handle("/protected", middleware.AuthMiddleware(http.HandlerFunc(handlers.ProtectedEndpoint)))

    fmt.Println("Server is running at :7070")
    log.Fatal(http.ListenAndServe(":7070", nil))
}
