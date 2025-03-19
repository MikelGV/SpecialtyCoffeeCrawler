package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type UserStore struct {
    Db *sql.DB
}

type User struct {
    Id int 
    Name string 
    Email string 
    Password string 
    CreatedAt string 
    UpdatedAt string 
}

/**
    Handles the Create user call to the database
**/
func (u UserStore) CreateUser(usr User) error {
    _, err := u.Db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", 
        usr.Name, usr.Email, usr.Password)

    if err != nil {
        return fmt.Errorf("error creating user: %w", err)
    }

    return nil
}

//func (u UserStore) GetAllUsers() ([]*User, error) {}

func (u UserStore) GetUsersById(id string) (*User, error) {
    var usr User

    err := u.Db.QueryRow("SELECT id, name, email FROM users WHER id = $1", id).Scan(
        &usr.Id, &usr.Name, &usr.Email)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, fmt.Errorf("error retrieving user: %w", err)
    }

    return &usr, nil
}

func (u UserStore) GetUsersByEmail(email string) (*User, error) {
    var usr User

    err := u.Db.QueryRow("SELECT id, name, email FROM users WHERE email = $1", email).Scan(
        &usr.Id, &usr.Name, &usr.Email)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }

        return nil, fmt.Errorf("error retrieving user: %w", err)
    }

    return &usr, nil
}

func (u UserStore) GetOrCreateUsersByEmail(email string) (*User, error) {
    var usr User

    err := u.Db.QueryRow("SELECT id, name, email FROM users WHERE email = $1", email).Scan(
        &usr.Id, &usr.Name, &usr.Email)

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            newUsr := User{Email: email, Password: ""}
            err := u.CreateUser(newUsr)

            if err != nil {
                return nil, fmt.Errorf("failed to create user: %w", err)
            }

            err = u.Db.QueryRow("SELECT id, name, email FROM users WHERE email = $1", email).Scan(
                &usr.Id, &usr.Name, &usr.Email)

            if err != nil {
                return nil, fmt.Errorf("failed to retrieve created user: %w", err)
            }

            return &newUsr, nil 
        }
        return nil, fmt.Errorf("error checking user: %w", err)
    }
    return &usr, nil
}

func (u UserStore) UpdateUser(usr User, userId string) (*User, error) {
    _, err := u.Db.Exec("UPDATE users SET (name, email, password, updated_at) WHERE (id) VALUES ($1, $2, $3, $4, $5)", 
        usr.Name, &usr.Email, &usr.Password, time.Now(), userId,
    )

    if err != nil {
        return nil, fmt.Errorf("error updating user: %w", err)
    }

    return u.GetUsersById(userId)
}

func (u UserStore) DeleteUser(id string) error {
    _, err := u.Db.Exec("DELETE users WHERE (id) VALUES ($1)", id)

    if err != nil {
        return fmt.Errorf("error deleting user: %w", err)
    }

    return nil
}

func scanRowsIntoUser(rows *sql.Rows) (*User, error) {
    usr := new(User)

    err := rows.Scan(
        &usr.Id,
        &usr.Name,
        &usr.Email,
        &usr.Password,
        &usr.CreatedAt,
        &usr.UpdatedAt,
    )

    if err != nil {
        return nil, fmt.Errorf("Unable to scan rows into user: %w", err)
    }

    return usr, nil
}
