package routes

import (
	"net/http"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/api"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/database"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/config"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
)

func AddRoutes(
    mux *http.ServeMux,
    conf config.Config,
    log *logger.Logger,
    usrStore *database.UserStore,
) {
    mux.Handle("/api/signup", api.PostCreateUserHandler(usrStore, log))
    mux.Handle("/api/login", api.LogInWithJWTHandler(log, usrStore))
    mux.Handle("/api/logout", api.LogoutHandler(log, usrStore))
    mux.Handle("/api/auth/google", api.Oauth2LoginHandler(log, usrStore))
    mux.Handle("/api/auth/google/callback", api.GoogleLoginCallback(log, usrStore))
}
