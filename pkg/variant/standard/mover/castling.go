package mover

import (
	"errors"

	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
)

// Castling это структура, реализующая интерфейс Mover для рокировки.
// Она отвечает за выполнение и проверку допустимости рокировки на доске.
type Castling struct {
}

func (m *Castling) Make(castlingType move.CastlingType, board chess.Board) (chess.Move, error) {
	if err := validator.ValidateCastling(castlingType, board); err != nil {
		return nil, err
	}

	direction := fileDirection(castlingType)

	_, kingPosition := board.Squares().GetPiece(chess.NotationKing, board.Turn())
	rookPosition, _ := m.rookPosition(direction, board.Squares(), kingPosition)

	rank := kingPosition.Rank

	board.MovePiece(kingPosition, position.New(kingPosition.File+direction*2, rank))
	board.MovePiece(rookPosition, position.New(kingPosition.File+direction, rank))

	board.NextTurn()
	defer board.NextTurn()

	return &move.Castling{
		CheckMate: &move.CheckMate{
			IsCheck: board.State().IsCheck(),
			IsMate:  board.State().IsMate(),
		},
		CastlingType: castlingType,
	}, nil
}

func (m *Castling) rookPosition(direction position.File, squares chess.Squares, kingPosition position.Position) (position.Position, error) {
	for i := kingPosition.File + direction; i <= position.FileH && i >= 0; i += direction {
		square := squares.GetByPosition(position.New(i, kingPosition.Rank))
		if !square.IsEmpty() && square.Piece.Notation() == chess.NotationRook {
			return square.Position, nil
		}
	}

	return position.Position{}, errors.New("ладья не найдена")
}

func fileDirection(castlingType move.CastlingType) position.File {
	return map[move.CastlingType]position.File{
		move.CastlingShort: 1,
		move.CastlingLong:  -1,
	}[castlingType]
}
