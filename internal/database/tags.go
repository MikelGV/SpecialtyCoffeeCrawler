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

func (t TagsStore) GetAllTags() ([]*Tags, error) {
    query := `SELECT * FROM tags`
    rows, err := t.DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("failed to query all tags: %v", err)
    }
    var allTags []*Tags

    for rows.Next() {
        roaster, err := scanRowsIntoTags(rows)
        if err != nil {
            return nil, fmt.Errorf("failed to scan tags: %v", err)
        }
        allTags = append(allTags, roaster)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }
    if len(allTags) == 0 {
        return nil, sql.ErrNoRows
    }

    return allTags, nil
}


func scanRowsIntoTags(rows *sql.Rows) (*Tags, error) {
    tgs := new(Tags)

    err := rows.Scan(
        &tgs.Id,
        &tgs.Name,
    )

    if err != nil {
        return nil, fmt.Errorf("Unable to scan rows into roasters: %w", err)
    }

    return tgs, nil
}
