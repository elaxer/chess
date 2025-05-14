package standard

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/mover"
)

// standard - эта структура описывает шахматную доску и ее состояние.
// Реализует логику стандартных шахмат.
// Реализует интерфейс standard из пакета chess.
type standard struct {
	turn           chess.Side
	squares        chess.Squares
	movesHistory   []chess.Move
	capturedPieces []chess.Piece
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

func (b *standard) State() chess.State {
	if b.isDraw() {
		return chess.StateDraw
	}

	isCheck := b.isCheck()
	availableMovesCount := b.Moves(b.turn).Cardinality()

	if isCheck && availableMovesCount == 0 {
		return chess.StateMate
	}
	if isCheck {
		return chess.StateCheck
	}
	if availableMovesCount == 0 {
		return chess.StateStalemate
	}

	return chess.StateClear
}

func (b *standard) MakeMove(move chess.Move) error {
	modifiedMove, err := mover.MakeMove(move, b)
	if err != nil {
		return err
	}

	b.movesHistory = append(b.movesHistory, modifiedMove)
	b.NextTurn()

	return nil
}

func (b *standard) NextTurn() {
	b.turn = !b.turn
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

func (b *standard) isCheck() bool {
	_, kingPosition := b.squares.GetPiece(chess.NotationKing, b.turn)

	return b.Moves(!b.turn).ContainsOne(kingPosition)
}

func (b *standard) isDraw() bool {
	return b.squares.GetAllPiecesCount(chess.SideWhite) == 1 &&
		b.squares.GetAllPiecesCount(chess.SideBlack) == 1 &&
		!b.isThreefoldRepetition()
}

func (b *standard) isThreefoldRepetition() bool {
	return false
}

// castlings возвращает возможные рокировки для текущей стороны.
// Если рокировка невозможна, то она не будет включена в список.
func (b *standard) castlings() []move.CastlingType {
	castlings := make([]move.CastlingType, 0, 2)

	if validator.ValidateCastling(move.CastlingShort, b) == nil {
		castlings = append(castlings, move.CastlingShort)
	}
	if validator.ValidateCastling(move.CastlingLong, b) == nil {
		castlings = append(castlings, move.CastlingLong)
	}

	return castlings
}

func (b *standard) String() string {
	return "todo: return FEN"
}

func (b *standard) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"squares":         b.squares,
		"state":           b.State(),
		"captured_pieces": b.capturedPieces,
		"castlings":       b.castlings(),
	})
}
