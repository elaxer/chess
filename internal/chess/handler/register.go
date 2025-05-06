package handler

import (
	"net/http"
	"text/template"

	"github.com/elaxer/chess/internal/chess/dto"
	"github.com/elaxer/chess/internal/chess/service"
)

type Register struct {
	rootDir string

	authService *service.Auth
	userService *service.User
}

func NewRegister(rootDir string, authService *service.Auth, userService *service.User) *Register {
	return &Register{rootDir, authService, userService}
}

// Register выдает страницу регистрации пользователя.
// Deprecated: use RegisterSubmit instead.
func (h *Register) Register(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles(h.rootDir+"/frontend/public/base.html", h.rootDir+"/frontend/public/register.html")).Execute(w, nil)
}

// RegisterSubmit обрабатывает запрос на регистрацию.
func (h *Register) RegisterSubmit(w http.ResponseWriter, r *http.Request) {
	registerReq, ok := ReadJSONBody[dto.RegisterRequest](w, r.Body)
	if !ok {
		return
	}

	if _, err := h.userService.Register(registerReq.Login, registerReq.Password); err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	token, err := h.authService.Authenticate(registerReq.Login, registerReq.Password)
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	ResponseJSON(w, dto.Tokens{Access: token}, http.StatusCreated)
}
