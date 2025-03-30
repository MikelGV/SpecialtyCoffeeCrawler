package database

import (
	"database/sql"
	"fmt"
)

type ProductsStore struct {
    DB *sql.DB
}

type Products struct {
    Id int
    Title string
    Price int
    Image string
    Origin string
    Type string
    Method string
    ProductUrl string
    CreateAT string
    UpdatedAt string

}

func (p ProductsStore) InsertProduct(pr Products, roasterId string) error {
    _, err := p.DB.Exec("INSERT INTO products (title, price, prImg, pUrl, type, origin, method, roaster_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
        pr.Title, pr.Price, pr.Image, pr.ProductUrl, pr.Origin, pr.Method, roasterId)

    if err != nil {
        return fmt.Errorf("error creating product: %w", err)
    }

    return nil
}

/**
func (p ProductsStore) GetAllProductsByRoasterID() (*Products, error) {
}

func (p ProductsStore) GetAllRoasterProductsByFilter() (*Products, error) {
}
**/
