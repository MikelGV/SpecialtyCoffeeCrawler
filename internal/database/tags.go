package database

import (
	"database/sql"
	"fmt"
)

type TagsStore struct {
    DB *sql.DB
}

type Tags struct {
    Id int
    Name string
}

func (t TagsStore) InsertTags(tagName string) error {
    _, err := t.DB.Exec("INSERT INTO tags (name) VALUES ($1)" )

    if err != nil {
        return fmt.Errorf("error creating tags: %w", err)
    }

    return nil

} 
