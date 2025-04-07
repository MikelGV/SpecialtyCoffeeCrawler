package api

import (
	"net/http"
	"strconv"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/utils"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/database"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
)

func CreateProductHandler(log *logger.Logger, product *database.ProductsStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            log.Error("Failed to get proper method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
    })
} 

func GetAllProductsByRoasterID(log * logger.Logger, product *database.ProductsStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet && r.Method != http.MethodOptions {
            log.Error("Failed to get proper method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        roasterId := r.URL.Query().Get("roasters_id")
        rId, err := strconv.Atoi(roasterId)
        if err != nil {
            log.Error("Failed to transform id into int", "error", err)
            http.Error(w, "Interal Server Error", http.StatusInternalServerError)
            return
        }

        type res struct {
            Id int `json:"product_id"`
            Title string `json:"title"`
            Origin string `json:"Origin"`
            Description string `json:"description"`
            Price float64 `json:"price"`
            WebUrl string `json:"web_url"`
            Method string `json:"method"`
            Type string `json:"type"`
            Image string `json:"pImg"`
            RoasterID string `json:"roaster_id"`
        }



        products, err := product.GetAllProductsByRoasterID(rId)
        if err != nil {
            log.Error("Failed to get roaster by id", "error", err)
            http.Error(w, "Interal Server Error", http.StatusInternalServerError)
            return
        }

        var resp []res
        for _, p := range products {
            resp = append(resp, res{
                Id: p.Id,
                Title: p.Title,
                Origin: p.Origin,
                Price: p.Price,
                WebUrl: p.ProductUrl,
                Method: p.Method,
                Type: p.Type,
                Image: p.Image,
                RoasterID: p.RoasterID,
            })
        }

        utils.Encode(w, r, http.StatusOK, resp)


    })
}

func GetAllRoasterProductsByFilters(log *logger.Logger, product *database.ProductsStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet && r.Method != http.MethodOptions {
            log.Error("Failed to get proper method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
    })
}
