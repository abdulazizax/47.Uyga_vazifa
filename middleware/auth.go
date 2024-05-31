package middleware

import (
    "context"
    "fmt"
    "net/http"
    "strings"
    "h47/utils"
)

type key int

const (
    roleKey key = iota
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Forbidden: No Authorization header", http.StatusForbidden)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := utils.ParseJWT(tokenString)
        if err != nil {
            fmt.Println("Error parsing JWT:", err)
            http.Error(w, "Forbidden: Invalid token", http.StatusForbidden)
            return
        }

        fmt.Println("Token valid. Role:", claims.Role)

        ctx := context.WithValue(r.Context(), roleKey, claims.Role)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
