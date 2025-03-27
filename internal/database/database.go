package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)


type DBStore struct {
    Conn *sql.DB
    Users *UserStore
    Tags *TagsStore
    UserTags *User_TagsStore
}

func Connect() (*DBStore, error) {
    db, err := sql.Open("pgx", "postgres://coffee:123456789@localhost:5432/specialty_coffee") 

    if err != nil {
        fmt.Fprintf(os.Stderr, "error opening database: %s\n", err)
        return nil, err 
    }

    if err := initUsers(db); err != nil {
        fmt.Fprintf(os.Stderr, "error creating users database table: %s\n", err)
        return nil, err 
    }

    if err := initTags(db); err != nil {
        fmt.Fprintf(os.Stderr, "error creating tags database table: %s\n", err)
        return nil, err
    }

    if err := initUser_Tags(db); err != nil {
        fmt.Fprintf(os.Stderr, "error creating user_tags database table: %s\n", err)
        return nil, err
    }

    if err := db.Ping(); err != nil {
        fmt.Fprintf(os.Stderr, "error pinging databasei: %s", err)
        return nil, err 
    }

    fmt.Println("Database Connection Stablished")
    return &DBStore{
        Conn: db,
        Users: &UserStore{db},
        Tags: &TagsStore{db},
        UserTags: &User_TagsStore{db},
    }, nil
}

/**
    Create User Table if it's not created
**/
func initUsers(db *sql.DB) error {
    query := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        role BOOL NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

    _, err := db.Exec(query)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating users table: %s", err)
        return err
    }

    fmt.Println("Users table is ready")
    return nil
}

func initTags(db *sql.DB) error {
    query := `CREATE TABLE IF NOT EXISTS tags (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL UNIQUE
    )` 

    _, err := db.Exec(query)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating tags table: %s", err)
        return err
    }

    fmt.Println("Tags table is ready")
    return nil
}

func initUser_Tags(db *sql.DB) error {
    query := `CREATE TABLE IF NOT EXISTS user_tags (
        user_id INT REFERENCES users(id),
        tag_id INT REFERENCES tags(id),
        PRIMARY KEY (user_id, tag_id)
    )`

    _, err := db.Exec(query)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating User_Tags table: %s", err)
        return err
    }

    fmt.Println("User_Tags table is ready")
    return nil
}
