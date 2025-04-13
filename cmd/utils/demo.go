package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/database"
)

type DemoRoasterRequest struct {
    Name string `json:"name"`
    Location string `json:"location"`
    Description string `json:"description"`
    WebsiteUrl string `json:"website_url"`
    ContactEmail string `json:"contact_email"`
}

type DemoProductRequest struct {
    Title string `json:"title"`
    Price float64 `json:"price"`
    Image string `json:"product_img"`
    Origin string `json:"origin"`
    Type string `json:"type"`
    Method string `json:"method"`
    ProductUrl string `json:"product_url"`
    RoasterID int `json:"roaster_id"`
}

func InitializeDemoDB(dbStore *database.DBStore) error {
    data, err := os.ReadFile("dummy_data.json")
    if err != nil {
        return fmt.Errorf("error reading dummy_data.json: %w", err)
    }

    var dummy struct {
        Roasters []DemoRoasterRequest `json:"roasters"`
        Products []DemoProductRequest `json:"products"`
    }

    if err := json.Unmarshal(data, &dummy); err != nil {
        if syntxErr, ok := err.(*json.SyntaxError); ok {
            fmt.Printf("syntax error at byte offest %d: %s\n", syntxErr.Offset, syntxErr.Error())
        }
        return fmt.Errorf("error unmarshaling dummy data: %w", err)
    }

    tx, err := dbStore.Conn.Begin()
    if err != nil {
        return fmt.Errorf("error starting transaction: %w", err)
    }
    
    defer tx.Rollback()

    roasterIDs := make(map[int]int)
    for i, r := range dummy.Roasters {
        query := `INSERT INTO roasters (name, location, description, websiteUrl, contact_email)
                    VALUES ($1, $2, $3, $4, $5) ON CONFLICT (name) DO UPDATE
                    SET location = EXCLUDED.location,
                        description = EXCLUDED.description,
                        websiteUrl = EXCLUDED.websiteUrl,
                        contact_email = EXCLUDED.contact_email,
                        updated_at = CURRENT_TIMESTAMP
                    RETURNING id`
        var id int
        err := tx.QueryRow(query, r.Name, r.Location, r.Description, r.WebsiteUrl, r.ContactEmail).Scan(&id)

        if err != nil {
            return fmt.Errorf("error inserting roaster %s: %w", r.Name, err)
        }
        roasterIDs[i+1] = id
    }

    for _, p := range dummy.Products {
        roasterID, ok := roasterIDs[p.RoasterID]

        if !ok {
            return fmt.Errorf("invalid roaster_id %d, for product %s", p.RoasterID, p.Title)
        }

        query := ` INSERT INTO products (title, price, prImg, pUrl, type, origin, method, roaster_id)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
            ON CONFLICT (title, roaster_id) DO UPDATE
            SET price = EXCLUDED.price,
                prImg = EXCLUDED.prImg,
                pUrl = EXCLUDED.pUrl,
                type = EXCLUDED.type,
                origin = EXCLUDED.origin,
                method = EXCLUDED.method,
                updated_at = CURRENT_TIMESTAMP
            RETURNING id
        `


        var productID int
        err := tx.QueryRow(query, p.Title, p.Price, p.Image, p.ProductUrl, p.Type, p.Origin, p.Method, roasterID).Scan(&productID)
        if err != nil {
            return fmt.Errorf("error inserting products %s: %w", p.Title, err)
        }

        roaster := dummy.Roasters[p.RoasterID-1]
        tags := []string{
            roaster.Location,
            p.Type,
            p.Method,
            p.Origin,
        }

        for _, tag := range tags {
            var tagID int
            tagQuery := `
                INSERT INTO tags (name)
                VALUES ($1) ON CONFLICT (name) DO NOTHING
                RETURNING id
            `

            err = tx.QueryRow(tagQuery, tag).Scan(&tagID)
            if err != nil && err != sql.ErrNoRows {
                return fmt.Errorf("error inserting tag %s: %w", tag, err)
            }

            if err == sql.ErrNoRows {
                err = tx.QueryRow("SELECT id FROM tags WHERE name =$1", tag).Scan(&tagID)
                if err != nil {
                    return fmt.Errorf("error fetching tag %s: %w", tag, err)
                }
            }

            linkQuery := `INSERT INTO product_tags (product_id, tag_id)
                            VALUES ($1, $2) ON CONFLICT DO NOTHING`
            _, err = tx.Exec(linkQuery, productID, tagID)
            if err != nil {
                return fmt.Errorf("error linking product %d to tag %d: %w", productID, tagID, err)
            }
        }

    }

    if err = tx.Commit(); err != nil {
        return fmt.Errorf("error commiting transaction: %w", err)
    }

    fmt.Println("Database initialized with dummy data from json")
    return nil
}
