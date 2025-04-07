package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/utils"
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

        roasterId := r.URL.Query().Get("id")
        rID, _ := strconv.Atoi(roasterId)
        
        roaster, err := roaster.GetRoasterById(rID)
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

func GetAllRoastersHandlers(log *logger.Logger, roaster *database.RoastersStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet && r.Method != http.MethodOptions {
            log.Error("Failed to get proper method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        type Response struct {
            Id int `json:"id"`
            Name string `json:"name"`
            Location string `json:"location"`
            Description string `json:"description"`
            WebsiteUrl string `json:"website_url"`
        }

        roasters, err := roaster.GetAllRoasters()
        if err != nil {
            log.Error("failed to get all roasters", "error", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return 
        }

        if len(roasters) == 0 {
            utils.Encode(w, r, http.StatusOK, []Response{})
        }

        var resp []Response
        for _, r := range roasters {
            resp = append(resp, Response{
                Id: r.Id,
                Name: r.Name,
                Location: r.Location,
                Description: r.Description,
                WebsiteUrl: r.WebsiteUrl,
            })
        }

        utils.Encode(w, r, http.StatusOK, resp)
    })
}

// This are temporary handlers to work with nextjs till we set up tailwind and templ
func GetRoasterHandler(log *logger.Logger, roaster *database.RoastersStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet && r.Method != http.MethodOptions {
            log.Error("Failed to get proper method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        roasterId := r.URL.Query().Get("id")
        rId, err := strconv.Atoi(roasterId)
        if err != nil {
            log.Error("Failed to transform id into int", "error", err)
            http.Error(w, "Interal Server Error", http.StatusInternalServerError)
            return
        }

        type res struct {
            Id int `json:"roaster_id"`
            Name string `json:"roaster_name"`
            Location string `json:"location"`
            Description string `json:"description"`
            WebsiteUrl string `json:"website_url"`
            ContactEmail string `json:"contact_email"`
        }



        roaster, err := roaster.GetRoasterById(rId)
        if err != nil {
            log.Error("Failed to get roaster by id", "error", err)
            http.Error(w, "Interal Server Error", http.StatusInternalServerError)
            return
        }

        resp := res{
            Id: roaster.Id,
            Name: roaster.Name,
            Location: roaster.Location,
            Description: roaster.Description,
            WebsiteUrl: roaster.WebsiteUrl,
            ContactEmail: roaster.ContactEmail,
        }

        utils.Encode(w, r, http.StatusOK, resp)
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
