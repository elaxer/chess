package handler

import (
	"html/template"
	"net/http"

	"github.com/elaxer/chess/internal/chess/dto"
	"github.com/elaxer/chess/internal/chess/service"
	"github.com/gorilla/sessions"
)

// todo
const sessionName = "session"

type Auth struct {
	rootDir string

	store sessions.Store

	authService *service.Auth
	userService *service.User
}

func NewAuth(rootDir string, store sessions.Store, authService *service.Auth, userService *service.User) *Auth {
	return &Auth{rootDir, store, authService, userService}
}

// Register выдает страницу регистрации пользователя.
func (h *Auth) Register(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles(h.rootDir+"/frontend/public/base.html", h.rootDir+"/frontend/public/register.html")).Execute(w, nil)
}

// RegisterSubmit обрабатывает запрос на регистрацию пользователя.
func (h *Auth) RegisterSubmit(w http.ResponseWriter, r *http.Request) {
	registerReq, ok := ReadJSONBody[dto.RegisterRequest](w, r.Body)
	if !ok {
		return
	}

	u, err := h.userService.Register(registerReq.Login, registerReq.Password)
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	token, err := h.authService.Authenticate(registerReq.Login, registerReq.Password)
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	session, _ := h.store.Get(r, sessionName)
	session.Values["user"] = u
	if err := session.Save(r, w); err != nil {
		ResponseError(w, err, http.StatusInternalServerError)
		return
	}

	ResponseJSON(w, dto.Tokens{Access: token}, http.StatusCreated)
}

// Login выдает страницу входа пользователя.
func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles(h.rootDir+"/frontend/public/base.html", h.rootDir+"/frontend/public/login.html")).Execute(w, nil)
}

// LoginSubmit обрабатывает запрос на вход пользователя.
func (h *Auth) LoginSubmit(w http.ResponseWriter, r *http.Request) {
	loginReq, ok := ReadJSONBody[dto.LoginRequest](w, r.Body)
	if !ok {
		return
	}

	token, err := h.authService.Authenticate(loginReq.Login, loginReq.Password)
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	ResponseJSON(w, dto.Tokens{Access: token}, http.StatusOK)
}
