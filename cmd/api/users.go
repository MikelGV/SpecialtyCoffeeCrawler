package api

import (
	"net/http"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/utils"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/cmd/web/templates"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/database"
	"github.com/MikelGV/SpecialtyCoffeeCrawler/internal/server/logger"
)

/**
    API endpoint that handles the Get Settings
**/

func GetUserSettingHandler(us *database.UserStore, log *logger.Logger) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            log.Error("Invalid Method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
        
        userId, err := utils.GetUserIdFromToken(r) 
        if err != nil {
            log.Error("couldn't get userId from token", "error", err)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        user, err := us.GetUsersById(userId)
        if err != nil {
            log.Error("couldn't get user from userID", "error", err)
            http.Error(w, "user not found", http.StatusInternalServerError)
            return
        }

        err = templates.SettingsPage(user).Render(r.Context(), w)
        if err != nil {
            log.Error("Tempolate rendering error", "error", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }
    })
}

/**
    API endpoint that handles the Post http request for creating a user.
**/
func PostCreateUserHandler(us *database.UserStore, log *logger.Logger) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            log.Error("Invalid Method", "method", r.Method)
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        type Request struct {
            Name string `json:"name"`
            Email string `json:"email"`
            Password string `json:"password"`
        }

        req, err := utils.Decode[Request](r)
        if  err != nil {
            log.Error("Failed to decode request", "error", err)
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        } 

        hashedPassword, err := utils.EncryptPassowrd(req.Password)

        if err != nil {
            log.Error("Failed to encrypt password", "error", err)
            http.Error(w, "internal server error", http.StatusInternalServerError)
            return
        }

        if err := us.CreateUser(database.User{
            Name: req.Name,
            Email: req.Email,
            Password: hashedPassword,
        }); err != nil {
            log.Error("Failed to create user", "error", err)
            http.Error(w, "internal server error", http.StatusInternalServerError)
            return
        }

        if err := utils.Encode(w, r, http.StatusOK, req); err != nil {
            log.Error("Failed to encode response", err)
            http.Error(w, "internal server error", http.StatusInternalServerError)
            return
        }

        w.Write([]byte("User created successfully!"))
    })
}

/**
    API endpoint that handles the PUT http request for updating a user
**/
func PutUpdateUserHandler(us *database.UserStore, log *logger.Logger) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPut {
            log.Error("Invalid Method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        //I need to add a GetUserIDFromToken function later
        userId, err := utils.GetUserIdFromToken(r)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

            
        type Request struct {
            Name string `json:"name"`
            Email string `json:"email"`
            Password string `json:"password"`
        }

        req, err := utils.Decode[Request](r);
        if  err != nil {
            log.Error("Failed to decode request", "error", err)
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        } 
        
        /**
            I belive here i should do something like redirect to 
            /profile/user.id or idk for now i will stay with _
        **/
         _, err = us.UpdateUser(database.User{
            Name: req.Name,
            Email: req.Email,
            Password: req.Password,
        }, userId)

        if err != nil {
            log.Error("Failed to update user", "error", err)
            http.Error(w, "Could not update user", http.StatusInternalServerError)
            return
        }

        utils.Encode(w, r, http.StatusOK, req)
        w.Write([]byte("User updated successfully!"))
    })
}

func DeleteUserHandler(us *database.UserStore, log *logger.Logger) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodDelete {
            log.Error("Failed to get the right method", "method", r.Method)
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return 
        }
    
        userId, err := utils.GetUserIdFromToken(r) 
        if err != nil {
            log.Error("couldn't get userId from token", "error", err)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        err = us.DeleteUser(userId)
        if err != nil {
            log.Error("couldn't get user from userID", "error", err)
            http.Error(w, "user not found", http.StatusInternalServerError)
            return
        }

        w.Write([]byte("User deleted successfully!"))

    })
} 
