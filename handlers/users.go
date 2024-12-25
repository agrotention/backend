package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/agrotention/backend/dto"
	"github.com/agrotention/backend/services"
	"github.com/agrotention/backend/utils"
)

// =========== Type User Handler
type UserHandler struct {
	svc *services.UserService
}

// =========== Constructor
func NewUserHandler(svc *services.UserService) *UserHandler {
	return &UserHandler{svc}
}

// =========== Register Handler to Mux
func (h *UserHandler) RegisterRouter(mux *http.ServeMux) {
	// Auth
	mux.HandleFunc("PUT /auth/change-email", h.handleChangeEmail)
	mux.HandleFunc("PUT /auth/change-password", h.handleChangePassword)
	mux.HandleFunc("POST /auth/login", h.handleLogin)
	mux.HandleFunc("POST /auth/signup", h.handleRegister)

	// User
	mux.HandleFunc("DELETE /users/{id}/delete", h.handleDelete)
	mux.HandleFunc("GET /users/{id}", h.handleDetail)
	mux.HandleFunc("DELETE /users/{id}", h.handleDisable)
	mux.HandleFunc("GET /users", h.handleList)
	mux.HandleFunc("PUT /users/{id}", h.handleUpdate)
}

// =========== Handlers
// Change user Email
func (h *UserHandler) handleChangeEmail(w http.ResponseWriter, r *http.Request) {
	var data dto.ReqUserChangeEmail
	res, err := h.svc.ChangeEmail(data)
	if err != nil {
		panic("unhandled")
	}
	json.NewEncoder(w).Encode(res)
}

// Change user password
func (h *UserHandler) handleChangePassword(w http.ResponseWriter, r *http.Request) {
	var data dto.ReqUserChangePassword
	res, err := h.svc.ChangePassword(data)
	if err != nil {
		panic("unhandled")
	}
	json.NewEncoder(w).Encode(res)
}

// Permanently delete user
func (h *UserHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	var data dto.ReqUserDelete
	res, err := h.svc.Delete(data)
	if err != nil {
		panic("unhandled")
	}
	json.NewEncoder(w).Encode(res)
}

// Get detail user
func (h *UserHandler) handleDetail(w http.ResponseWriter, r *http.Request) {
	var data dto.ReqUserDetail
	res, err := h.svc.Detail(data)
	if err != nil {
		panic("unhandled")
	}
	json.NewEncoder(w).Encode(res)
}

// Disable user
func (h *UserHandler) handleDisable(w http.ResponseWriter, r *http.Request) {
	var data dto.ReqUserDisable
	res, err := h.svc.Disable(data)
	if err != nil {
		panic("unhandled")
	}
	json.NewEncoder(w).Encode(res)
}

// List users
func (h *UserHandler) handleList(w http.ResponseWriter, r *http.Request) {
	var data dto.ReqUserList
	res, err := h.svc.List(data)
	if err != nil {
		panic("unhandled")
	}
	json.NewEncoder(w).Encode(res)
}

// Login user
func (h *UserHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var data dto.ReqUserLogin
	res, err := h.svc.Login(data)
	if err != nil {
		panic("unhandled")
	}
	json.NewEncoder(w).Encode(res)
}

// Signup (register) new user
func (h *UserHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// Read JSON data
	var data dto.ReqUserRegister
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.LogErr.Println(err)
		utils.ErrInternal.Send(w)
		return
	}

	// Send to service
	if res, err := h.svc.Register(data); err != nil {
		err.Send(w)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(res)
	}

}

// Update user info
func (h *UserHandler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	var data dto.ReqUserUpdate
	res, err := h.svc.Update(data)
	if err != nil {
		panic("unhandled")
	}
	json.NewEncoder(w).Encode(res)
}
