package api

import (
	"net/http"

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

func GetRoasterProfileHandler(log *logger.Logger, roaster *database.RoastersStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            log.Error("Failed to get proper method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
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
