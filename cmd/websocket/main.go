package main

import (
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/elaxer/chess/internal/chess/repository"
	"github.com/elaxer/chess/internal/chess/service"
	"github.com/elaxer/chess/internal/config"
	"github.com/elaxer/chess/internal/database"
	"github.com/elaxer/chess/internal/websocket/handler"
	"github.com/gorilla/websocket"
	"github.com/huandu/go-sqlbuilder"
	"github.com/joho/godotenv"
)

var (
	_, b, _, _ = runtime.Caller(0)
	rootDir, _ = filepath.Abs(filepath.Dir(b) + "/../..")
)

func init() {
	must(godotenv.Load(rootDir + "/.env"))

	sqlbuilder.DefaultFlavor = config.SQLFlavor
}

func main() {
	cfg := config.MustLoad()

	db := database.MustConnectAndPing(cfg.DB)

	gameRepository := repository.NewGameDB(db)
	userRepository := repository.NewUserDB(db)

	gameService := service.NewGame(gameRepository)
	authService := service.NewAuth(userRepository)

	websocketHandler := handler.NewWebsocket(
		websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		gameService,
		authService,
	)

	http.Handle("/ws", websocketHandler)

	must(http.ListenAndServe(":8080", nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
