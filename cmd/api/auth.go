package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/utils"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/database"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/config"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

//var SecretKey = []byte(config.Env.JwtSecretKey)

var googleOauthConfig = &oauth2.Config {
    ClientID: config.Env.GClientID,
    ClientSecret: config.Env.GCSecret,
    RedirectURL: config.Env.GRedirectUrl,
    Scopes: []string{config.Env.GScopeUrl},
    Endpoint: google.Endpoint,
}

/**
    JWT login Handler
**/
func LogInWithJWTHandler(log *logger.Logger, user *database.UserStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost  && r.Method != http.MethodOptions{
            log.Error("Method Not Allowed", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        type Request struct {
            Email string `json:"email"`
            Password string `json:"password"`
        }

        type Response struct {
            Message string `json:"message"`
            Token string `json:"token,omitempty"`
        }

        req, err := utils.Decode[Request](r)
        if err != nil {
            log.Error("failed to decode", "error", err)
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        }

        user, err := user.GetUsersByEmail(req.Email)
        if err != nil {
            http.Error(w, "Invalid email or password", http.StatusUnauthorized)
            return
        }

        token, err := createToken(user.Id)
        if err != nil {
            log.Error("error creating a token", "error", err)
            http.Error(w, "internal server error", http.StatusInternalServerError)
            return
        }

        http.SetCookie(w, &http.Cookie{
            Name: "auth_token",
            Value: token,
            HttpOnly: true,
            Expires: time.Now().Add(time.Hour * 24),
            Path: "/",
        })

        res := Response{
            Message: "User logged in successfully!",
            Token: token,
        }

        utils.Encode(w, r, http.StatusOK, res)
    })
}

/**
    OAuth2 login Handler and Callback 
**/

func Oauth2LoginHandler(log *logger.Logger, user *database.UserStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        url := googleOauthConfig.AuthCodeURL("randomstate")
        http.Redirect(w, r, url, http.StatusTemporaryRedirect)
    })
}

func GoogleLoginCallback(log *logger.Logger, user *database.UserStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            log.Error("Method Not Allowed", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        code := r.URL.Query().Get("code")
        if code == "" {
            log.Error("code missing")
            http.Error(w, "Missing Code", http.StatusBadRequest)
            return
        }

        token, err := googleOauthConfig.Exchange(context.Background(), code)
        if err != nil {
            log.Error("failed to get token", "error", err)
            http.Error(w, "Failed to get token", http.StatusInternalServerError)
            return
        }

        client := googleOauthConfig.Client(context.Background(), token)

        resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
        if err != nil {
            log.Error("failed to get user info", "error", err)
            http.Error(w, "Failed to get user info", http.StatusInternalServerError)
            return
        }

        defer resp.Body.Close()

        type userInfo struct {
            Email string `json:"email"`
        }

        info, err := utils.Decode[userInfo](r)
        if err != nil {
            log.Error("Failed to decode request", "error", err)
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }

        user, err := user.GetOrCreateUsersByEmail(info.Email)

        if err != nil {
            log.Error("error geting or creating user", "error", err)
            http.Error(w, "failed to get or create user", http.StatusInternalServerError)
            return
        }

        jwtToken, err := createToken(user.Id)
        if err != nil {
            log.Error("error creating a token", "error", err)
            http.Error(w, "internal server error", http.StatusInternalServerError)
            return
        }

        http.SetCookie(w, &http.Cookie{
            Name: "auth_token",
            Value: jwtToken,
            HttpOnly: true,
            Expires: time.Now().Add(time.Hour * 24),
            Path: "/",
        })

        http.Redirect(w, r, "/", http.StatusFound)
    })
}

/**
    Lgout Handler
**/

func LogoutHandler(log *logger.Logger, user *database.UserStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            log.Error("Method Not Allowed", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        http.SetCookie(w, &http.Cookie{
            Name: "auth_token",
            Value: "",
            Expires: time.Now().Add(-1 * time.Hour),
            Path: "/",
        })
        log.Info("user logged to the home page (or login page)")

        http.Redirect(w, r, "/", http.StatusFound)
    })
}

/**
    Here go helper functions
**/

func createToken(userId int) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userId": userId,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString(utils.SecretKey)

    if err != nil {
        return "", fmt.Errorf("Token couldn't be created: %w", err)
    }

    return tokenString, nil
}



