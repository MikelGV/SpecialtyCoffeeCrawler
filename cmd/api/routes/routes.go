package routes

import (
	"net/http"

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
    fs := http.FileServer(http.Dir("/cmd/web/assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

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
    mux.Handle("/api/roasters", api.GetAllRoastersHandlers(log, rstStore))
    mux.Handle("/api/roasters/id", api.GetRoasterHandler(log, rstStore))
    mux.Handle("/api/products", api.GetAllProductsByRoasterID(log, productStore))
    mux.Handle("/api/products/filters", middleware.AuthMiddleware(api.GetAllRoasterProductsByFilters(log, productStore)))

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
