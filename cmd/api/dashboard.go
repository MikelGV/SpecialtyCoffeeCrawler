package api

import (
	"net/http"

//    "github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/web/assets/layout"
    "github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/web"

)

func GetHomePage() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
