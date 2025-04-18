package config

import (
	"os"
	"strconv"

)

type Config struct {
    Host string
    Port string
    JwtSecretKey string
    GClientID string
    GCSecret string
    GScopeUrl string
    GRedirectUrl string
    DB_URL string
}


var Env = initConfig()

func initConfig() Config {
    return Config{
        Host: getEnv("HOST", "localhost"),
        Port: getEnv("PORT", "3000"),
        JwtSecretKey: getEnv("JWTSECRETKEY", "somethingSuperSecret"),
        GClientID: getEnv("GCCLIENTID", "failed"),
        GCSecret: getEnv("GCSECRET", "somethingSecret"),
        GScopeUrl: getEnv("GSOPEURL", "someScope"),
        GRedirectUrl: getEnv("GREDIRECTURL", "someUrl"),
        DB_URL: getEnv("DB_URL", ""),
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
