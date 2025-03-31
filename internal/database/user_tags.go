package database

import (
	"database/sql"
	"fmt"
)

type User_TagsStore struct {
    DB *sql.DB
}

type UserTags struct {
    UserID int
    TagsID int
}

func (us User_TagsStore) InsertUserTags(userId, tagId string) error {
    _, err := us.DB.Exec("INSERT INTO user_tags (user_id, tag_id) VALUES ($1, $2)", userId, tagId)

    if err != nil {
        return fmt.Errorf("error creating product: %w", err)
    }

    return nil
}

func (t User_TagsStore) GetUserTags(userId string) ([]string, error) {
    query := `SELECT t.name FROM user_tags ut
    JOIN tags t ON t.id = ut.tag_id
    WHERE ut.user_id = $1 ORDER BY t.name`

    rows, err := t.DB.Query(query, userId) 
    if err != nil {
        return nil, fmt.Errorf("failed to query user tags: %w", err)
    }
    
    defer rows.Close()

    var tagNames []string
    for rows.Next() {
        var tagName string
        if err := rows.Scan(*&tagName); err != nil {
            return nil, fmt.Errorf("failed to scan user tag: %w", err)
        }
        tagNames = append(tagNames, tagName)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %v", err)
    }


    if len(tagNames) == 0 {
        return nil, sql.ErrNoRows
    }

    return tagNames, nil

}

