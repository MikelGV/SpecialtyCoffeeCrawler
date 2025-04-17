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
    JwtSecretKey string
    GClientID string
    GCSecret string
    GScopeUrl string
    GRedirectUrl string
    DB_Password string
    DB_Database string
    DB_Username string
    DB_Port string
}


var Env = initConfig()

func initConfig() Config {
    err := godotenv.Load()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error loading godotenv %s\n", err)
    }

    return Config{
        Host: getEnv("HOST", "localhost"),
        Port: getEnv("PORT", "3000"),
        JwtSecretKey: getEnv("JWTSECRETKEY", "somethingSuperSecret"),
        GClientID: getEnv("GCCLIENTID", "failed"),
        GCSecret: getEnv("GCSECRET", "somethingSecret"),
        GScopeUrl: getEnv("GSOPEURL", "someScope"),
        GRedirectUrl: getEnv("GREDIRECTURL", "someUrl"),
        DB_Password: getEnv("DB_PASSWORD", "123456789"),
        DB_Username: getEnv("DB_USERNAME", "coffee"),
        DB_Database: getEnv("DB_DATABASE", "specialty_coffee"),
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
