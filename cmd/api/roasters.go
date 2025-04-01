package api

import (
	"net/http"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/web/templates"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/database"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
)

type RoasterRequest struct {
    Name string `json:"name"`
    Location string `json:"location"`
    Description string `json:"description"`
    WebsiteUrl string `json:"website_url"`
    ContactEmail string `json:"contact_email"`
}

/**
    Get the roasters profile by the id
**/
func GetRoasterProfileHandler(log *logger.Logger, roaster *database.RoastersStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            log.Error("Failed to get proper method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        roasterId := r.URL.Query().Get("roaster_id")
        
        roaster, err := roaster.GetRoasterById(roasterId)
        if err != nil {
            log.Error("failed to fetch roaster by id", "error", err  )
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }

        w.Header().Set("Content-Type", "text/html")
        err = templates.RoasertProfile(roaster).Render(r.Context(), w)
        if err != nil {
            log.Error("error to render roaster profile", "error", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
    })
} 

func CreateRoasterHandler(log *logger.Logger, roaster *database.RoastersStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            log.Error("Failed to get proper method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
    })
} 
