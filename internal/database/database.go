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

    if err := db.Ping(); err != nil {
        fmt.Fprintf(os.Stderr, "error pinging databasei: %s", err)
        return nil, err 
    }

    fmt.Println("Database Connection Stablished")
    return &DBStore{
        Conn: db,
        Users: &UserStore{db},
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
        password TEXT NOT NULL UNIQUE,
        role TEXT NOT NULL,
        tags TEXT,
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
