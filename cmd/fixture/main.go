package main

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/elaxer/chess/internal/chess/repository"
	"github.com/elaxer/chess/internal/config"
	"github.com/elaxer/chess/internal/database"
	"github.com/elaxer/chess/internal/fixture"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/jackc/pgx/v5/stdlib"
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

	userRepository := repository.NewUserDB(db)
	gameRepository := repository.NewGameDB(db)

	tx, ctx := database.MustBeginTx(context.Background(), db)

	defer func() {
		if r := recover(); r != nil {
			must(tx.Rollback())
			panic(r)
		}
	}()

	for _, u := range fixture.Users {
		must(userRepository.Add(ctx, u))
	}
	for _, g := range fixture.Games {
		must(gameRepository.Add(ctx, g))
	}

	must(tx.Commit())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
