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
}
