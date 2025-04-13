package database

import (
	"database/sql"
	"fmt"
)

type ProductTagsStore struct {
    DB *sql.DB
}

type ProductTags struct {
    ProductID int
    TagsID int
}

func (p ProductTagsStore) InserProductTags(producId, tagID string) error {
    _, err := p.DB.Exec("INSERT INTO product_tags (product_id, tag_id) VALUES ($1, $2)", producId, tagID)

    if err != nil {
        return fmt.Errorf("error creating product: %w", err)
    }

    return nil
}
