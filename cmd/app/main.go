package main

import (
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/elaxer/chess/internal/chess/handler"
	"github.com/elaxer/chess/internal/chess/handler/middleware"
	"github.com/elaxer/chess/internal/chess/repository"
	"github.com/elaxer/chess/internal/chess/service"
	"github.com/elaxer/chess/internal/config"
	"github.com/elaxer/chess/internal/database"
	"github.com/gorilla/sessions"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var (
	_, b, _, _ = runtime.Caller(0)
	rootDir, _ = filepath.Abs(filepath.Dir(b) + "/../..")

	mux = http.NewServeMux()
)

func init() {
	must(godotenv.Load(rootDir + "/.env"))

	sqlbuilder.DefaultFlavor = config.SQLFlavor
}

func main() {
	cfg := config.MustLoad()

	db := database.MustConnectAndPing(cfg.DB)
	// todo
	store := sessions.NewCookieStore([]byte("secret_key"))

	userR := repository.NewUserDB(db)
	gameR := repository.NewGameDB(db)

	authS := service.NewAuth(userR)
	userS := service.NewUser(userR)

	authH := handler.NewAuth(rootDir, store, authS, userS)
	gameH := handler.NewGame(rootDir, db, userR, gameR)

	fs := http.FileServer(http.Dir(rootDir + config.PublicDir))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))

	mux.Handle("/games/{id}", middleware.Auth(http.HandlerFunc(gameH.Game), authS))
	mux.Handle("/", middleware.Auth(http.HandlerFunc(gameH.List), authS))
	mux.Handle("/users/{id}", middleware.Auth(http.HandlerFunc(gameH.ListByUser), authS))

	mux.Handle("/register", middleware.Guest(http.HandlerFunc(authH.Register)))
	mux.Handle("/register_submit", middleware.Guest(middleware.Method(http.HandlerFunc(authH.RegisterSubmit), http.MethodPost)))
	mux.Handle("/login", middleware.Guest(http.HandlerFunc(authH.Login)))
	mux.Handle("/login_submit", middleware.Guest(middleware.Method(http.HandlerFunc(authH.LoginSubmit), http.MethodPost)))

	must(http.ListenAndServe(":8081", mux))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
