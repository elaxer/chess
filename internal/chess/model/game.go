package model

import (
	"slices"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/pgstr"
	"github.com/google/uuid"
)

// Game - это модель игры.
// Содержит идентификаторы игроков, доску и историю ходов.
type Game struct {
	*BaseModel
	// PlayerWhiteID - это идентификатор игрока, который играет белыми.
	PlayerWhiteID uuid.UUID `db:"player_white_id"`
	// PlayerBlackID - это идентификатор игрока, который играет черными.
	PlayerBlackID uuid.UUID `db:"player_black_id"`

	PlayerWhite *User
	PlayerBlack *User

	Board    chess.Board
	MovesStr string `db:"moves"`
	// Moves - это история ходов в алгебраической нотации.
	Moves []string
	// Result - это исход игры.
	Result GameResult `db:"result"`
}

func NewGame(board chess.Board, playerWhiteID, playerBlackID uuid.UUID) *Game {
	return &Game{
		BaseModel:     newBaseModel(),
		PlayerWhiteID: playerWhiteID,
		PlayerBlackID: playerBlackID,
		Board:         board,
		Result:        GameResultInProcess,
	}
}

func NewGameExisted(playerWhiteID, playerBlackID uuid.UUID, result GameResult, moves []string) *Game {
	return &Game{
		BaseModel:     newBaseModel(),
		PlayerWhiteID: playerWhiteID,
		PlayerBlackID: playerBlackID,
		Moves:         moves,
		Result:        result,
	}
}

func NewGameEmpty(playerWhiteID, playerBlackID uuid.UUID) *Game {
	// return NewGame(standard.NewBoard(), playerWhiteID, playerBlackID)
	return nil
}

// AddMove добавляет ход в историю игры.
func (g *Game) AddMove(moveNotation string, moveNotations ...string) {
	g.Moves = slices.Concat(g.Moves, []string{moveNotation}, moveNotations)
}

// End завершает игру, записывая исход игры.
func (g *Game) End(lastMove chess.Move) {
	// todo
	// if lastMove.MoveType.IsDraw() || lastMove.MoveType.IsStalemate() {
	// 	g.Result = GameResultDraw
	// 	return
	// }

	// if !lastMove.MoveType.IsMate() {
	// 	return
	// }

	// if g.Board.Turn.IsBlack() {
	// 	g.Result = GameResultWinWhite
	// 	return
	// }

	g.Result = GameResultWinBlack
}

// Leave завершает игру, если один из игроков покинул игру.
func (g *Game) Leave(leavedPlayerID uuid.UUID) {
	switch leavedPlayerID {
	case g.PlayerWhiteID:
		g.Result = GameResultWinBlack
	case g.PlayerBlackID:
		g.Result = GameResultWinWhite
	default:
		return
	}

	g.Result |= GameResultLeaved
}

func (g *Game) AfterGet() {
	g.Moves = pgstr.Parse(g.MovesStr)
	g.MovesStr = ""
}
