package standard

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/mover"
	"github.com/elaxer/chess/pkg/variant/standard/staterule"
)

// standard - эта структура описывает шахматную доску и ее состояние.
// Реализует логику стандартных шахмат.
// Реализует интерфейс standard из пакета chess.
type standard struct {
	turn           chess.Side
	squares        chess.Squares
	movesHistory   []chess.Move
	capturedPieces []chess.Piece

	stateRules []staterule.Rule
}

func (b *standard) Squares() chess.Squares {
	return b.squares
}

func (b *standard) Turn() chess.Side {
	return b.turn
}

func (b *standard) MovesHistory() []chess.Move {
	return b.movesHistory
}

func (b *standard) Moves(side chess.Side) position.Set {
	moves := mapset.NewSetWithSize[position.Position](32)
	for _, piece := range b.squares.GetAllPieces(side) {
		moves = moves.Union(piece.Moves(b))
	}

	return moves
}

func (b *standard) State(side chess.Side) chess.State {
	for _, rule := range b.stateRules {
		if state := rule(b, side); state != chess.StateClear {
			return state
		}
	}

	return chess.StateClear
}

func (b *standard) MakeMove(move chess.Move) error {
	modifiedMove, err := mover.MakeMove(move, b)
	if err != nil {
		return err
	}

	b.movesHistory = append(b.movesHistory, modifiedMove)
	b.turn = !b.turn

	return nil
}

func (b *standard) MovePiece(from, to position.Position) (capturedPiece chess.Piece) {
	fromSquare := b.squares.GetByPosition(from)
	fromSquare.Piece.SetMoved()
	defer fromSquare.SetPiece(nil)

	toSquare := b.squares.GetByPosition(to)
	defer toSquare.SetPiece(fromSquare.Piece)

	if capturedPiece = toSquare.Piece; capturedPiece != nil {
		b.capturedPieces = append(b.capturedPieces, capturedPiece)
	}

	return
}

// castlings возвращает возможные рокировки для текущей стороны.
// Если рокировка невозможна, то она не будет включена в список.
func (b *standard) castlings() []move.CastlingType {
	castlings := make([]move.CastlingType, 0, 2)

	if validator.ValidateCastling(move.CastlingShort, b.turn, b) == nil {
		castlings = append(castlings, move.CastlingShort)
	}
	if validator.ValidateCastling(move.CastlingLong, b.turn, b) == nil {
		castlings = append(castlings, move.CastlingLong)
	}

	return castlings
}

func (b *standard) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"squares":         b.squares,
		"state":           b.State(b.turn),
		"captured_pieces": b.capturedPieces,
		"castlings":       b.castlings(),
	})
}
