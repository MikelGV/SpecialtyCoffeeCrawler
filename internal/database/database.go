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
    Products *ProductsStore
    Roasters *RoastersStore
    ProductTags *ProductTagsStore
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

    if err := initRoasters(db); err != nil {
        fmt.Fprintf(os.Stderr, "error creating roasters database table: %s\n", err)
        return nil, err
    }

    if err := initProducts(db); err != nil {
        fmt.Fprintf(os.Stderr, "error creating products database table: %s\n", err)
        return nil, err
    }

    if err := initProduct_Tags(db); err != nil {
        fmt.Fprintf(os.Stderr, "error creating products_tags database table: %s\n", err)
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
        Products: &ProductsStore{db},
        Roasters: &RoastersStore{db},
        ProductTags: &ProductTagsStore{db},
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
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

    _, err := db.Exec(query)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating users table: %s", err)
        return err
    }

    // This has to be removed after demo is over
    insertUserQuery := `INSERT INTO users (name, email, password) 
                        VALUES ($1, $2, $3) ON CONFLICT (email) DO NOTHING`

    _, err = db.Exec(insertUserQuery, "TestUser", "test@test.com", "thisisatest")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating users: %s", err)
        return err
    }

    // This stays
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

// Maybe
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

func initProducts(db *sql.DB) error {
    query := `CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        price DECIMAL(10,2) NOT NULL,
        prImg TEXT NOT NULL,
        pUrl TEXT NOT NULL,
        type TEXT NOT NULL,
        origin TEXT NOT NULL,
        method TEXT NOT NULL, 
        roaster_id INT NOT NULL REFERENCES roasters(id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT unique_title_roaster UNIQUE (title, roaster_id)
    )`

    _, err := db.Exec(query)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating products table: %s", err)
        return err
    }

    fmt.Println("Products table is ready")
    return nil
}

func initProduct_Tags(db *sql.DB) error {
    query := `CREATE TABLE IF NOT EXISTS product_tags (
        product_id INT REFERENCES products(id),
        tag_id INT REFERENCES tags(id),
        PRIMARY KEY (product_id, tag_id)
    )`

    _, err := db.Exec(query)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating Product_Tags table: %s", err)
        return err
    }

    fmt.Println("Product_Tags table is ready")
    return nil
}

func initRoasters(db *sql.DB) error {
    query := `CREATE TABLE IF NOT EXISTS roasters (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL UNIQUE,
        location TEXT NOT NULL,
        description TEXT,
        websiteUrl TEXT NOT NULL,
        contact_email TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

    _, err := db.Exec(query)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating roasters table: %s", err)
        return err 
    }

    fmt.Println("Roasters table is ready")
    return nil
}
