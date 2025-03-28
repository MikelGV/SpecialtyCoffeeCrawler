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
