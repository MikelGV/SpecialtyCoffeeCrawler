package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/config"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)
var SecretKey = []byte(config.Env.JwtSecretKey)

func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    if err := json.NewEncoder(w).Encode(v); err != nil {
        return fmt.Errorf("encode json: %w", err)
    }
    return nil
}

func Decode[T any](r *http.Request) (T, error) {
    var v T

    if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
        return v, fmt.Errorf("decode json: %w", err)
    }

    return v, nil
}

func EncryptPassowrd(password string) (string, error) {

    encrypt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    
    if err != nil {
        logger.NewLogger().Error("Failed to encrypt password", err)
        return "", fmt.Errorf("error hashsing password: %w", err)
    }

    return string(encrypt), nil
}

func ComparePassword(encrypted string, pWord []byte) bool {
    err := bcrypt.CompareHashAndPassword([]byte(encrypted), pWord)
    return err == nil 
}


func GetUserIdFromToken(r *http.Request) (string, error) {
    cookie, err := r.Cookie("auth_token")
    if err != nil {
        return "", errors.New("missing auth token")
    }

    token, err := jwt.Parse(cookie.Value, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return SecretKey, nil
    })

    if err != nil || !token.Valid {
        return "", errors.New("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)

    if !ok {
        return "", errors.New("invalid token claims")
    }

    userId, ok := claims["userId"].(string)
    if !ok {
        return "", errors.New("user ID not found in token")
    }

    return userId, nil 
}
