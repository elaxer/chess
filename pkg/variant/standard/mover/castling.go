package mover

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	validator "github.com/elaxer/chess/pkg/variant/standard/movevalidator"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state"
)

// Castling это структура, реализующая интерфейс Mover для рокировки.
// Она отвечает за выполнение и проверку допустимости рокировки на доске.
type Castling struct {
}

func (m *Castling) Make(castlingType move.CastlingType, board chess.Board) (chess.Move, error) {
	if err := validator.ValidateCastling(castlingType, board.Turn(), board, true); err != nil {
		return nil, err
	}

	direction := fileDirection(castlingType)

	_, kingPosition := board.Squares().GetPiece(piece.NotationKing, board.Turn())
	rookPosition, _ := m.rookPosition(direction, board.Squares(), kingPosition)

	rank := kingPosition.Rank

	board.Squares().MovePiece(kingPosition, position.New(kingPosition.File+direction*2, rank))
	board.Squares().MovePiece(rookPosition, position.New(kingPosition.File+direction, rank))

	return &move.Castling{
		CheckMate: &move.CheckMate{
			IsCheck: board.State(!board.Turn()) == state.Check,
			IsMate:  board.State(!board.Turn()) == state.Mate,
		},
		CastlingType: castlingType,
	}, nil
}

func (m *Castling) rookPosition(direction position.File, squares *chess.Squares, kingPosition position.Position) (position.Position, error) {
	for position, p := range squares.IterByDirection(kingPosition, position.New(direction, 0)) {
		if p != nil && p.Notation() == piece.NotationRook {
			return position, nil
		}
	}

	return position.NewNull(), fmt.Errorf("%w: rook wasn't found", validator.ErrCastling)
}

// todo
func fileDirection(castlingType move.CastlingType) position.File {
	return map[move.CastlingType]position.File{
		move.CastlingShort: 1,
		move.CastlingLong:  -1,
	}[castlingType]
}
