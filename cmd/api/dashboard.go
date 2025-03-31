package api

import (
	"net/http"

	//    "github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/web/assets/layout"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/utils"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/web"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/web/templates"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/database"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
)

func GetHomePage(log *logger.Logger) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            log.Error("method mismatch", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }

        err := web.HomePage().Render(r.Context(), w)
        if err != nil {
            http.Error(w, "Error rendering home", http.StatusInternalServerError)
        }
    })
}

func GetDashboard(log *logger.Logger, roaster *database.RoastersStore, user_tags *database.User_TagsStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            log.Error("method mismatch", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        userId, err := utils.GetUserIdFromToken(r) 
        if err != nil {
            log.Error("couldn't get userId from token", "error", err)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        userTags, err := user_tags.GetUserTags(userId)
        if err != nil {
            log.Error("failed to fetch user tags", "error", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return 
        }


        var roasters []*database.Roasters
        if len(userTags) > 0 {
            roasters, err = roaster.GetAllRoastersByUser_Tags(userTags)
        } else {
            roasters, err = roaster.GetAllRoasters()
        }
        if err != nil {
            log.Error("failed to fetch all roasters", "error", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/html")
        err = templates.DashboardPage(roasters).Render(r.Context(), w)
        if err != nil {
            log.Error("failed to render dashboard template", "error", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
    })
}
