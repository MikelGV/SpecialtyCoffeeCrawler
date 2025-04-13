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
    Price float64
    Image string
    Origin string
    Type string
    Method string
    ProductUrl string
    RoasterID string
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

func (p ProductsStore) GetAllProductsByRoasterID(roaster_id int) ([]*Products, error) {
    query := `SELECT id, title, price, prImg, pUrl, type, origin, method, roaster_id FROM products WHERE roaster_id = $1 ORDER BY title`

    rows, err := p.DB.Query(query, roaster_id)

    if err != nil {
        return nil, fmt.Errorf("failed to query roasters by product tags: %v", err)
    }

    defer rows.Close()

    var allProducts []*Products

    for rows.Next() {
        products, err := scanRowsIntoProducts(rows)
        if err != nil {
            return nil, fmt.Errorf("failed to scan roaster: %v", err)
        }
        allProducts  = append(allProducts, products)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }
    if len(allProducts) == 0 {
        return nil, sql.ErrNoRows
    }

    return allProducts, nil
}

func (p ProductsStore) GetAllRoasterProductsByFilter(tagNames []string) ([]*Products, error) {
    placeholders := make([]string, len(tagNames))
    args := make([]interface{}, len(tagNames))
    for i, tag := range tagNames {
        placeholders[i] = fmt.Sprintf("$%d", i+1) 
        args[i] = tag
    }

    query := `SELECT DISTINCT p.id, p.title, p.price, p.prImg, p.pUrl, p.type, p.method, p.origin, p.roaster_id 
        FROM products p
        JOIN roasters r ON r.id = p.roaster_id 
        JOIN product_tags pt ON pt.product_id = p.id
        JOIN tags t ON t.id = pt.tag_id
        WHERE t.name IN ($1) ORDER BY r.name
    `

    rows, err := p.DB.Query(query, args...)

    if err != nil {
        return nil, fmt.Errorf("failed to query roaster products by tags: %v", err)
    }

    defer rows.Close()

    var allProducts []*Products

    for rows.Next() {
        products, err := scanRowsIntoProducts(rows)
        if err != nil {
            return nil, fmt.Errorf("failed to scan roaster: %v", err)
        }
        allProducts = append(allProducts, products)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }
    if len(allProducts) == 0 {
        return nil, sql.ErrNoRows
    }

    return allProducts, nil

}

func (p ProductsStore) GetAllProductsByUser_Tags(roaster_id int, tagNames []string) ([]*Products, error) {
    placeholders := make([]string, len(tagNames))
    args := make([]interface{}, len(tagNames))
    args = append(args, roaster_id)
    for i, tag := range tagNames {
        placeholders[i] = fmt.Sprintf("$%d", i+2) 
        args[i] = tag
    }

    query := `SELECT DISTINCT p.id, p.title, p.price, p.prImg, p.pUrl, p.type, p.method, p.origin, p.roaster_id
        FROM products p
        JOIN user_tags pu ON pu.product_id = p.id
        JOIN tags t ON t.id = pu.tag_id
        WHERE p.roaster_id ($1) AND t.name IN ($2) ORDER BY r.name
    `

    rows, err := p.DB.Query(query, args...)

    if err != nil {
        return nil, fmt.Errorf("failed to query roaster products by tags: %v", err)
    }

    defer rows.Close()

    var allProducts []*Products

    for rows.Next() {
        products, err := scanRowsIntoProducts(rows)
        if err != nil {
            return nil, fmt.Errorf("failed to scan products: %v", err)
        }
        allProducts = append(allProducts, products)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }
    if len(allProducts) == 0 {
        return nil, sql.ErrNoRows
    }

    return allProducts, nil

}

func scanRowsIntoProducts(rows *sql.Rows) (*Products, error) {
    pr := new(Products)

    err := rows.Scan(
        &pr.Id,
        &pr.Title,
        &pr.Price, 
        &pr.Image,
        &pr.Origin,
        &pr.Type,
        &pr.Method,
        &pr.ProductUrl,
        &pr.CreateAT,
    )

    if err != nil {
        return nil, fmt.Errorf("Unable to scan rows into products: %w", err)
    }

    return pr, nil
}
