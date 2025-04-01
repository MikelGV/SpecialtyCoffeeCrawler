package routes

import (
	"net/http"
	"strings"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/api"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/middleware"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/database"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/config"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
)

func AddRoutes(
    mux *http.ServeMux,
    conf config.Config,
    log *logger.Logger,
    usrStore *database.UserStore,
    rstStore *database.RoastersStore,
    productStore *database.ProductsStore,
    productTagsStore *database.ProductTagsStore,
    tagsStore *database.TagsStore,
    usrTagStore *database.User_TagsStore,
) {
    fs := http.FileServer(http.Dir("cmd/web/assets"))
    mux.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
        path := strings.TrimPrefix(r.URL.Path, "/assets")

        if strings.HasPrefix(path, "/") {
            path = path[1:]
        }

        if strings.HasSuffix(path, ".css") {
            w.Header().Set("Content-Type", "text/css; charset=utf-8")
        } else if strings.HasSuffix(path, ".js") {
            w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
        } else {
            w.Header().Set("Content-Type", "application/octet-stream")
        }

        http.StripPrefix("/assets/", fs).ServeHTTP(w, r)
    })

    /**
        Logged Out pages
    **/
    mux.Handle("/", api.GetHomePage(log))
    mux.Handle("/signup", api.GetSignUp(log))
    mux.Handle("/login", api.GetLogIn(log))

    /**
        Loged In Pages
    **/
    mux.Handle("/dashboard", middleware.AuthMiddleware(api.GetDashboard(log, rstStore, usrTagStore)))
    mux.Handle("/roaster", middleware.AuthMiddleware(api.GetRoasterProfileHandler(log, rstStore)))

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
    mux.Handle("/settings", middleware.AuthMiddleware(api.GetUserSettingHandler(usrStore, log)))
    mux.Handle("/api/settings/update", middleware.AuthMiddleware(api.PutUpdateUserHandler(usrStore, log)))
    mux.Handle("/api/settings/delete", middleware.AuthMiddleware(api.DeleteUserHandler(usrStore, log)))
}
