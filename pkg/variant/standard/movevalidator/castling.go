package validator

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

var ErrCastling = fmt.Errorf("%w: castling validation error", Err)

func ValidateCastling(castlingType move.CastlingType, side chess.Side, board chess.Board, validateObstacle bool) error {
	king, kingPosition := board.Squares().GetPiece(piece.NotationKing, side)
	if king == nil {
		return fmt.Errorf("%w: the king wasn't found", ErrCastling)
	}
	if king.IsMoved() {
		return fmt.Errorf("%w: the king already has been moved", ErrCastling)
	}
	if !board.State(side).Type().IsClear() {
		return fmt.Errorf("%w: the king is under threat", ErrCastling)
	}

	rook, err := board.Squares().GetByPosition(castlingRookPosition(castlingType, kingPosition.Rank))
	if err != nil {
		return err
	}
	if rook == nil {
		return fmt.Errorf("%w: the rook wasn't found", ErrCastling)
	}
	if rook.IsMoved() {
		return fmt.Errorf("%w: the rook already has been moved", ErrCastling)
	}

	direction := fileDirection(castlingType)

	if validateObstacle {
		if err := castlingValidateObstacle(direction, board.Squares(), kingPosition, rook); err != nil {
			return err
		}
	}

	positions := mapset.NewSet(
		position.New(kingPosition.File+direction, kingPosition.Rank),
		position.New(kingPosition.File+direction*2, kingPosition.Rank),
	)
	if board.Moves(!side).Intersect(positions).Cardinality() > 0 {
		return fmt.Errorf("%w: castling squares are under threat", ErrCastling)
	}

	return nil
}

func castlingValidateObstacle(direction position.File, squares *chess.Squares, kingPosition position.Position, castlingRook chess.Piece) error {
	for _, piece := range squares.IterByDirection(kingPosition, position.New(direction, 0)) {
		if piece != nil && piece != castlingRook {
			return fmt.Errorf("%w: an obstacle")
		}
	}

	return nil
}

func fileDirection(castlingType move.CastlingType) position.File {
	return map[move.CastlingType]position.File{
		move.CastlingShort: 1,
		move.CastlingLong:  -1,
	}[castlingType]
}

func castlingRookPosition(castlingType move.CastlingType, rank position.Rank) position.Position {
	if castlingType == move.CastlingShort {
		return position.New(position.FileH, rank)
	} else {
		return position.New(position.FileA, rank)
	}
}
