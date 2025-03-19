package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
	"golang.org/x/crypto/bcrypt"
)

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
