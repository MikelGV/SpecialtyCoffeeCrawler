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
    /**
        This are the login, signup, and logout routes
    **/
    mux.Handle("/api/signup", api.PostCreateUserHandler(usrStore, log))
    mux.Handle("/api/login", api.LogInWithJWTHandler(log, usrStore))
    mux.Handle("/api/logout", api.LogoutHandler(log, usrStore))
    mux.Handle("/api/auth/google", api.Oauth2LoginHandler(log, usrStore))
    mux.Handle("/api/auth/google/callback", api.GoogleLoginCallback(log, usrStore))

    /**
        This are the routes that handle all of the user manipulation
        AKA: user profiles, update user, and delete user
    **/
    mux.Handle("/api/settings", api.GetUserSettingHandler(usrStore, log))
    mux.Handle("/api/settings/update", api.PutUpdateUserHandler(usrStore, log))
    mux.Handle("/api/settings/delete", api.DeleteUserHandler(usrStore, log))
}
