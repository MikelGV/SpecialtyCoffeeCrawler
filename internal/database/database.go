package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)


type DBStore struct {
    Conn *sql.DB
}

func Connect() (*DBStore, error) {
    db, err := sql.Open("pgx", "postgres://coffee:123456789@localhost:5432/specialty_coffee") 

    if err != nil {
        fmt.Fprintf(os.Stderr, "error opening database: %s\n", err)
        return nil, err 
    }

    if err := db.Ping(); err != nil {
        fmt.Fprintf(os.Stderr, "error pinging databasei: %s", err)
        return nil, err 
    }

    fmt.Println("Database Connection Stablished")
    return &DBStore{
        Conn: db,
    }, nil
}

/**
    Create User Table if it's not created
**/
func initUsers(db *sql.DB) error {
}
