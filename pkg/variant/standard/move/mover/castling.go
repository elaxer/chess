package mover

import (
	"fmt"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/result"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

// Castling это структура, реализующая интерфейс Mover для рокировки.
// Она отвечает за выполнение и проверку допустимости рокировки на доске.
type Castling struct {
}

func (m *Castling) Make(castlingType move.Castling, board chess.Board) (chess.MoveResult, error) {
	if err := validator.ValidateCastlingMove(castlingType, board.Turn(), board, true); err != nil {
		return nil, err
	}

	direction := fileDirection(castlingType)

	_, kingPosition := board.Squares().FindPiece(piece.NotationKing, board.Turn())
	rookPosition, _ := m.rookPosition(direction, board.Squares(), kingPosition)

	rank := kingPosition.Rank

	board.Squares().MovePiece(kingPosition, position.New(kingPosition.File+direction*2, rank))
	board.Squares().MovePiece(rookPosition, position.New(kingPosition.File+direction, rank))

	return &result.Castling{Abstract: newAbstractResult(board), Castling: castlingType}, nil
}

// todo
func (m *Castling) rookPosition(direction position.File, squares *chess.Squares, kingPosition position.Position) (position.Position, error) {
	for position, p := range squares.IterByDirection(kingPosition, position.New(direction, 0)) {
		if p != nil && p.Notation() == piece.NotationRook {
			return position, nil
		}
	}

	return position.NewEmpty(), fmt.Errorf("%w: rook wasn't found", validator.ErrCastling)
}

// todo
func fileDirection(castlingType move.Castling) position.File {
	return map[move.Castling]position.File{
		move.CastlingShort: 1,
		move.CastlingLong:  -1,
	}[castlingType]
}
