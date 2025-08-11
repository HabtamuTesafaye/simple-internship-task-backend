package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/minab/internship-backend/internal/model"
	"github.com/minab/internship-backend/internal/service"
	"github.com/minab/internship-backend/internal/util"
)

type UserHandler struct {
	service *service.UserService
}

type UserResponse struct {
	ID          string    `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Get a user by ID
// @Description Get a single user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} UserResponse
// @Failure 400 {string} string "Missing user ID"
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
// @Security BearerAuth
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	id := parts[len(parts)-1]
	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// @Summary Create a new user
// @Description Create a user with the given data
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body model.CreateUserRequest true "User Data"
// @Success 201 {object} UserResponse
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Failed to create user"
// @Router /register [post]
// @Security BearerAuth
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdUser, err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Map to response struct without password
	resp := UserResponse{
		ID:          createdUser.ID,
		FullName:    createdUser.FullName,
		Email:       createdUser.Email,
		PhoneNumber: createdUser.PhoneNumber,
		Role:        createdUser.Role,
		CreatedAt:   createdUser.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// @Summary Update a user
// @Description Update fields of a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param updates body map[string]interface{} true "Fields to update"
// @Success 200 {object} UserResponse
// @Failure 400 {string} string "Invalid request body or missing user ID"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Failed to update user"
// @Router /users/{id} [put]
// @Security BearerAuth
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}
	id := parts[len(parts)-1]

	// Fetch the existing user
	existing, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the request body into a map
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update only provided fields
	if v, ok := updates["full_name"].(string); ok {
		existing.FullName = v
	}
	if v, ok := updates["email"].(string); ok {
		existing.Email = v
	}
	if v, ok := updates["phone_number"].(string); ok {
		existing.PhoneNumber = v
	}
	if v, ok := updates["role"].(string); ok {
		existing.Role = v
	}
	if v, ok := updates["password"].(string); ok && v != "" {
		hashed, err := util.HashPassword(v)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		existing.Password = hashed
	}

	// Save the updated user
	if _, err := h.service.UpdateUser(r.Context(), id, existing); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Fetch the updated user from DB to ensure all fields are fresh
	updated, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to fetch updated user", http.StatusInternalServerError)
		return
	}

	// Map to response struct (do NOT include password)
	resp := UserResponse{
		ID:          updated.ID,
		FullName:    updated.FullName,
		Email:       updated.Email,
		PhoneNumber: updated.PhoneNumber,
		Role:        updated.Role,
		CreatedAt:   updated.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// @Summary List all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} UserResponse
// @Failure 500 {string} string "Failed to list users"
// @Router /users [get]
// @Security BearerAuth
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers(r.Context())
	if err != nil {
		http.Error(w, "Failed to list users", http.StatusInternalServerError)
		return
	}
	var resp []UserResponse
	for _, u := range users {
		resp = append(resp, UserResponse{
			ID:          u.ID,
			FullName:    u.FullName,
			Email:       u.Email,
			PhoneNumber: u.PhoneNumber,
			Role:        u.Role,
			CreatedAt:   u.CreatedAt,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
