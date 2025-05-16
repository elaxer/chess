package validator

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

var ErrCastling = fmt.Errorf("%w: ошибка валидации рокировки", Err)

func ValidateCastling(castlingType move.CastlingType, side chess.Side, board chess.Board) error {
	king, kingPosition := board.Squares().GetPiece(piece.NotationKing, side)
	if king == nil {
		return fmt.Errorf("%w: король не найден", ErrCastling)
	}
	if king.IsMoved() {
		return fmt.Errorf("%w: король уже ходил", ErrCastling)
	}
	if !board.State(side).IsClear() {
		return fmt.Errorf("%w: король под угрозой", ErrCastling)
	}

	positions, err := castlingVerifyingPositions(fileDirection(castlingType), board.Squares(), kingPosition)
	if err != nil {
		return err
	}

	if board.Moves(!side).Intersect(positions).Cardinality() > 0 {
		return fmt.Errorf("%w: поле для рокировки под боем", ErrCastling)
	}

	return nil
}

func castlingVerifyingPositions(direction position.File, squares chess.Squares, kingPosition position.Position) (position.Set, error) {
	positions := mapset.NewSetWithSize[position.Position](2)
	for file := kingPosition.File + direction; file <= squares.EdgePosition().File && file > 0; file += direction {
		pos := position.New(file, kingPosition.Rank)

		p, err := squares.GetByPosition(pos)
		if err != nil {
			return nil, fmt.Errorf("%w: нет ладьи", ErrCastling)
		}

		if p == nil {
			if diff := file - kingPosition.File; max(diff, -diff) <= 2 {
				positions.Add(pos)
			}

			continue
		}

		if p.Notation() != piece.NotationRook {
			return nil, fmt.Errorf("%w: помеха для рокировки", ErrCastling)
		}
		if p.IsMoved() {
			return nil, fmt.Errorf("%w: ладья уже ходила", ErrCastling)
		}

		break
	}

	return positions, nil
}

func fileDirection(castlingType move.CastlingType) position.File {
	return map[move.CastlingType]position.File{
		move.CastlingShort: 1,
		move.CastlingLong:  -1,
	}[castlingType]
}
