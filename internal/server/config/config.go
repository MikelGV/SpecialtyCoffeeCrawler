package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
    Host string
    Port string
}


var env = initConfig()

func initConfig() Config {
    err := godotenv.Load()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error loading godotenv %s\n", err)
    }

    return Config{
        Host: getEnv("HOST", "localhost"),
        Port: getEnv("PORT", "3000"),
    }
}


func getEnv(key, fallback string) string {
    if val, ok := os.LookupEnv(key); ok {
        return val
    }
    return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
    if val, ok := os.LookupEnv(key); ok {
        i, err := strconv.ParseInt(val, 10, 64)
        if err != nil {
            return fallback
        }
        return i
    }
    return fallback
}
