package standard

import (
	"encoding/json"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/set"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/mover"
)

// board структура описывает шахматную доску и ее состояние.
// Реализует логику стандартных шахмат.
// Реализует интерфейс board из пакета chess.
type board struct {
	turn           chess.Side
	squares        chess.Squares
	movesHistory   []chess.Move
	capturedPieces []chess.Piece
}

func (b *board) Squares() chess.Squares {
	return b.squares
}

func (b *board) Turn() chess.Side {
	return b.turn
}

func (b *board) MovesHistory() []chess.Move {
	return b.movesHistory
}

func (b *board) Moves(side chess.Side) *position.Set {
	moves := set.FromSlice(make([]position.Position, 0, 32))
	for _, piece := range b.squares.GetAllPieces(side) {
		moves = moves.Union(piece.Moves(b))
	}

	return moves
}

func (b *board) State() chess.State {
	if b.isDraw() {
		return chess.StateDraw
	}

	isCheck := b.isCheck(b.turn)
	availableMovesCount := b.Moves(b.turn).Len()

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

func (b *board) MakeMove(move chess.Move) error {
	modifiedMove, err := mover.MakeMove(move, b)
	if err != nil {
		return err
	}

	b.movesHistory = append(b.movesHistory, modifiedMove)
	b.NextTurn()

	return nil
}

func (b *board) UndoMove() (chess.Move, error) {
	if len(b.movesHistory) == 0 {
		return nil, nil
	}

	lastMove := b.movesHistory[len(b.movesHistory)-1]

	if err := mover.Undo(lastMove, b); err != nil {
		return nil, err
	}

	return lastMove, nil
}

func (b *board) NextTurn() {
	b.turn = !b.turn
}

func (b *board) MovePiece(from, to position.Position) (capturedPiece chess.Piece) {
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

func (b *board) isCheck(side chess.Side) bool {
	_, kingPosition := b.squares.GetPiece(chess.NotationKing, side)

	return b.Moves(!side).Has(kingPosition)
}

func (b *board) isDraw() bool {
	return b.squares.GetAllPiecesCount(chess.SideWhite) == 1 &&
		b.squares.GetAllPiecesCount(chess.SideBlack) == 1 &&
		!b.isThreefoldRepetition()
}

func (b *board) isThreefoldRepetition() bool {
	return false
}

// castlings возвращает возможные рокировки для текущей стороны.
// Если рокировка невозможна, то она не будет включена в список.
func (b *board) castlings() []move.CastlingType {
	castlings := make([]move.CastlingType, 0, 2)

	if err := validator.ValidateCastling(move.CastlingShort, b); err == nil {
		castlings = append(castlings, move.CastlingShort)
	}
	if err := validator.ValidateCastling(move.CastlingLong, b); err == nil {
		castlings = append(castlings, move.CastlingLong)
	}

	return castlings
}

func (b *board) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"squares":         b.squares,
		"state":           b.State(),
		"captured_pieces": b.capturedPieces,
		"castlings":       b.castlings(),
	})
}
