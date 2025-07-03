package middleware

import (
    "net/http"
    "strings"
    "user-service/utils"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
            return
        }

        token := strings.TrimPrefix(authHeader, "Bearer ")
        _, err := utils.VerifyJWT(token)
        if err != nil {
            http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    }
}
