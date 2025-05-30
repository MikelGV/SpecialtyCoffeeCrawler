package middleware

import (
	"net/http"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/config"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = config.Env.JwtSecretKey

/** Thi is for templ/htmx
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("auth_token")
        if err != nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return 
        }

        token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
            return SecretKey, nil
        })

        if err != nil || !token.Valid {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        next.ServeHTTP(w, r)
    })
}
**/
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("auth_token")
        if err != nil {
            if err == http.ErrNoCookie {
                w.Header().Set("Content-Type", "application/json")
                http.Error(w, `{"error": "No authentication cookie provided"}`, http.StatusUnauthorized)
                return 
            }
            w.Header().Set("Content-Type", "application/json")
            http.Error(w, `{"error": "Internal server error while checking cookie"}`, http.StatusInternalServerError)
            return 
        }

        token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
            return SecretKey, nil
        })

        if err != nil || !token.Valid {
                w.Header().Set("Content-Type", "application/json")
                http.Error(w, `{"error": "No valid token provided"}`, http.StatusUnauthorized)
                return 
        }

        next.ServeHTTP(w, r)
    })
}
