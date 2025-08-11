package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/minab/internship-backend/internal/service"
	"github.com/minab/internship-backend/internal/util"
)

type PasswordResetHandler struct {
	service *service.PasswordResetService
}

func NewPasswordResetHandler(service *service.PasswordResetService) *PasswordResetHandler {
	return &PasswordResetHandler{service: service}
}

// @Summary Request password reset
// @Description Send password reset link to user's email
// @Tags password
// @Accept  json
// @Produce  json
// @Param email body map[string]string true "User email"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "User not found"
// @Router /forgot-password [post]
func (h *PasswordResetHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	token, err := h.service.GenerateToken(r.Context(), req.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Build the reset link
	resetLink := os.Getenv("FRONTEND_URL") + "/reset-password?token=" + token

	// Load and render the HTML template
	tmpl, err := template.ParseFiles("internal/templates/reset_password.html")
	if err != nil {
		http.Error(w, "Failed to load email template", http.StatusInternalServerError)
		return
	}
	var bodyBuf bytes.Buffer
	if err := tmpl.Execute(&bodyBuf, map[string]string{"ResetLink": resetLink}); err != nil {
		http.Error(w, "Failed to render email template", http.StatusInternalServerError)
		return
	}

	// Send the email
	if err := util.SendEmail(req.Email, "Reset Your Password", bodyBuf.String()); err != nil {
		log.Printf("Failed to send email to %s: %v", req.Email, err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset link sent"})
}

// @Summary Reset password
// @Description Reset user password using token
// @Tags password
// @Accept  json
// @Produce  json
// @Param reset body map[string]string true "Token and new password"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid request or expired token"
// @Router /reset-password [post]
func (h *PasswordResetHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := h.service.ResetPassword(r.Context(), req.Token, req.NewPassword); err != nil {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password updated"})
}
