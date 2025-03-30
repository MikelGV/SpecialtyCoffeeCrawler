package api

import (
	"net/http"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/database"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
)

type ProductsRequest struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Price int `json:"price"`
    Image string `json:"product_img"`
    Origin string `json:"origin"`
    Type string `json:"type"`
    Method string `json:"method"`
    ProductUrl string `json:"product_url"`
    RoasterID int `json:"roaster_id"`
}

func CreateProductHandler(log *logger.Logger, product *database.ProductsStore) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            log.Error("Failed to get proper method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
    })
} 
