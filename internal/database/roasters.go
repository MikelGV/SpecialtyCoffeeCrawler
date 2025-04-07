package database

import (
	"database/sql"
	"fmt"
)

type RoastersStore struct {
    DB *sql.DB
}

type Roasters struct {
    Id int
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

func (r RoastersStore) GetRoasterById(roaster_id int) (*Roasters, error) {
    var rst Roasters

    err := r.DB.QueryRow("SELECT id, name, location, description, websiteUrl, contact_email FROM roasters WHERE id = $1", roaster_id).Scan(
        &rst.Id, &rst.Name, &rst.Location, &rst.Description, &rst.WebsiteUrl, &rst.ContactEmail)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("error retrieving roaster: %w", err)
    }

    return &rst, nil
}

func (r RoastersStore) GetAllRoasters() ([]*Roasters, error) {
    query := `SELECT * FROM roasters`
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to query all roasters: %v", err)
    }
    var allRoasters []*Roasters

    for rows.Next() {
        roaster, err := scanRowsIntoRoasters(rows)
        if err != nil {
            return nil, fmt.Errorf("failed to scan roaster: %v", err)
        }
        allRoasters = append(allRoasters, roaster)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }
    if len(allRoasters) == 0 {
        return nil, sql.ErrNoRows
    }

    return allRoasters, nil

}

func (r RoastersStore) GetAllRoasterByProductsFilter(tagNames []string) ([]*Roasters, error) {

    placeholders := make([]string, len(tagNames))
    args := make([]interface{}, len(tagNames))
    for i, tag := range tagNames {
        placeholders[i] = fmt.Sprintf("$%d", i+1) 
        args[i] = tag
    }

    query := `SELECT DISTINCT r.id, r.name, r.location, r.description, r.websiteUrl, r.contact_email
        FROM roasters r
        JOIN products ON p.roaster_id = r.id
        JOIN product_tags pt ON pt.product_id = p.id
        JOIN tags t ON t.id = pt.tag_id
        WHERE t.name IN ($1) ORDER BY r.name
    `

    rows, err := r.DB.Query(query, args...)

    if err != nil {
        return nil, fmt.Errorf("failed to query roasters by product tags: %v", err)
    }

    defer rows.Close()

    var allRoasters []*Roasters

    for rows.Next() {
        roaster, err := scanRowsIntoRoasters(rows)
        if err != nil {
            return nil, fmt.Errorf("failed to scan roaster: %v", err)
        }
        allRoasters = append(allRoasters, roaster)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }
    if len(allRoasters) == 0 {
        return nil, sql.ErrNoRows
    }

    return allRoasters, nil
}


func (r RoastersStore) GetAllRoastersByLocation(location string) ([]*Roasters, error) {
    query := `SELECT id, name, email FROM users WHERE location = $1`

    rows, err := r.DB.Query(query, location)
    if err != nil {
        return nil, fmt.Errorf("error querying roasters by location: %w", err)
    }

    var allRoasters []*Roasters

    for rows.Next() {
        roaster, err := scanRowsIntoRoasters(rows)
        if err != nil {
            return nil, fmt.Errorf("failed to scan roaster: %v", err)
        }
        allRoasters = append(allRoasters, roaster)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }
    if len(allRoasters) == 0 {
        return nil, sql.ErrNoRows
    }

    return allRoasters, nil
}


func (r RoastersStore) GetAllRoastersByUser_Tags(tagNames []string) ([]*Roasters, error) {
    placeholders := make([]string, len(tagNames))
    args := make([]interface{}, len(tagNames))
    for i, tag := range tagNames {
        placeholders[i] = fmt.Sprintf("$%d", i+2) 
        args[i] = tag
    }

    query := `SELECT DISTINCT r.id, r.name, r.location, r.description, r.websiteUrl, r.contact_email
        FROM roasters r
        JOIN products p ON p.roaster_id = r.id
        JOIN user_tags pu ON pu.product_id = p.id
        JOIN tags t ON t.id = pu.tag_id
        WHERE p.roaster_id ($1) AND t.name IN ($2) ORDER BY r.name
    `

    rows, err := r.DB.Query(query, args...)

    if err != nil {
        return nil, fmt.Errorf("failed to query roaster products by tags: %v", err)
    }

    defer rows.Close()

    var allRoasters []*Roasters

    for rows.Next() {
        roaster, err := scanRowsIntoRoasters(rows)
        if err != nil {
            return nil, fmt.Errorf("failed to scan products: %v", err)
        }
        allRoasters = append(allRoasters, roaster)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }
    if len(allRoasters) == 0 {
        return nil, sql.ErrNoRows
    }

    return allRoasters, nil

}
/**
**/

func scanRowsIntoRoasters(rows *sql.Rows) (*Roasters, error) {
    rst := new(Roasters)

    err := rows.Scan(
        &rst.Id,
        &rst.Name,
        &rst.Location,
        &rst.Description,
        &rst.WebsiteUrl,
        &rst.ContactEmail,
        &rst.CreatedAt,
        &rst.UpdatedAT,
    )

    if err != nil {
        return nil, fmt.Errorf("Unable to scan rows into roasters: %w", err)
    }

    return rst, nil
}
