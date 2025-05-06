package service

import (
	"context"

	"github.com/elaxer/chess/internal/chess/model"
	"github.com/elaxer/chess/internal/chess/repository"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/google/uuid"
)

type Game struct {
	gameRepository repository.Game
}

func NewGame(gameRepository repository.Game) *Game {
	return &Game{gameRepository}
}

// MakeMove делает ход в игре.
// Если ход является завершающим, то игра сохраняется в репозитории.
func (m *Game) MakeMove(game *model.Game, move chess.Move) error {
	//todo
	// if err := game.Board.MakeMove(move); err != nil {
	// 	return err
	// }

	// game.AddMove(fmt.Sprintf("%s", move))

	// if move.MoveType.IsGameOver() {
	// 	game.End(move)

	// 	return m.gameRepository.Add(context.Background(), game)
	// }

	return nil
}

func (m *Game) Leave(game *model.Game, playerID uuid.UUID) error {
	game.Leave(playerID)

	return m.gameRepository.Add(context.Background(), game)
}
