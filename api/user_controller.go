package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nanoohlaing1997/online-ordering-items/models"
)

type SignUpUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Address  string `json:"address" validate:"required"`
}

type SignInUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenUser struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"refresh_token" validate:"required"`
}

func (c *Controller) HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Ordering Item Online Management!!!"))
}

func (c *Controller) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var signUpUser SignUpUser

	if err := json.NewDecoder(r.Body).Decode(&signUpUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use Validator to perform struct validation
	if err := validate.Struct(signUpUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check usename and email already exit or not
	if _, err := c.dbm.GetUser(signUpUser.Email); err == nil {
		http.Error(w, "User already exit!!! Please login", http.StatusConflict)
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := HashPassword(signUpUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusFailedDependency)
		return
	}

	// Create JWT token
	jwtToken, err := GenerateJWT(signUpUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store hashed password, jwt token and refresh token to simulated user database
	user := &models.User{
		Name:         signUpUser.Username,
		Email:        signUpUser.Email,
		Password:     hashedPassword,
		Address:      signUpUser.Address,
		Status:       int32(0),
		RefreshToken: refreshToken,
	}
	if _, err := c.dbm.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Response with token
	response := &AuthResponse{
		Token:        jwtToken,
		RefreshToken: refreshToken,
	}
	json.NewEncoder(w).Encode(response)
}

func (c *Controller) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var signInUser SignInUser

	if err := json.NewDecoder(r.Body).Decode(&signInUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use Validator to perform struct validation
	if err := validate.Struct(signInUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := c.dbm.GetUser(signInUser.Email)
	if user == nil {
		http.Error(w, "User doesn't exist!!! Please sign up first", http.StatusNotFound)
		return
	}

	result := VerifyPassword(user.Password, signInUser.Password)
	if result {
		// Create JWT token
		jwtToken, err := GenerateJWT(user.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := &AuthResponse{
			Token:        jwtToken,
			RefreshToken: user.RefreshToken,
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	http.Error(w, "User unauthorized", http.StatusNotAcceptable)
}

func (c *Controller) TokenRefreshHandler(w http.ResponseWriter, r *http.Request) {
	var tokenUser RefreshTokenUser

	if err := json.NewDecoder(r.Body).Decode(&tokenUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Use Validator to perform struct validation
	if err := validate.Struct(tokenUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, _ := c.dbm.GetUser(tokenUser.Email)
	if user == nil {
		http.Error(w, "User doesn't exist!!! Please sign up first", http.StatusNotFound)
		return
	}

	fmt.Println(user.RefreshToken)
	if user.RefreshToken != tokenUser.Token {
		http.Error(w, "Token Unauthorized!!! Please sign up first", http.StatusUnauthorized)
		return
	}

	// Create JWT token
	jwtToken, err := GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &AuthResponse{
		Token:        jwtToken,
		RefreshToken: user.RefreshToken,
	}
	json.NewEncoder(w).Encode(response)
}
