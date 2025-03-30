package database

import (
	"database/sql"
	"fmt"
)

type RoastersStore struct {
    DB *sql.DB
}

type Roasters struct {
    Name string
    Location string
    Description string
    WebsiteUrl string
    ContactEmail string
    CreatedAt string
    UpdatedAT string
}

func (r RoastersStore)InsertRoaster(rst Roasters) error {
    _, err := r.DB.Exec("INSERT INTO roasters (name, location, description, websiteUrl, contact_email) VALUES ($1, $2, $3, $4, $5)", 
        rst.Name, rst.Location, rst.Description, rst.WebsiteUrl, rst.ContactEmail)

    if err != nil {
        return fmt.Errorf("error creating roaster: %w", err)
    }

    return nil
}

/**
func (r RoastersStore) GetAllRoasters() ([]*Roasters, error) {
}

func (r RoastersStore) GetAllRoasterByProductsFilter() ([]*Roasters, error) {}

func (r RoastersStore) GetAllRoastersByLocation() ([]*Roasters, error) {}

func (r RoastersStore) GetAllRoastersByLocationAndProductTags() ([]*Roasters, error) {}
**/
