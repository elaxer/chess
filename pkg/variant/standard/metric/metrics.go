package metric

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/metric"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/move"
	validator "github.com/elaxer/chess/pkg/variant/standard/movevalidator"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

var AllFuncs = []metric.MetricFunc{
	Castlings,
	CastlingsString,
}

func Castlings(board chess.Board) metric.Metric {
	castlings := make([]move.CastlingType, 0, 2)
	if validator.ValidateCastling(move.CastlingShort, board.Turn(), board) == nil {
		castlings = append(castlings, move.CastlingShort)
	}
	if validator.ValidateCastling(move.CastlingLong, board.Turn(), board) == nil {
		castlings = append(castlings, move.CastlingLong)
	}

	return metric.New("Available castlings", castlings)
}

func CastlingsString(board chess.Board) metric.Metric {
	castlings := map[string]bool{
		"K": validator.ValidateCastling(move.CastlingShort, chess.SideWhite, board) == nil,
		"Q": validator.ValidateCastling(move.CastlingLong, chess.SideWhite, board) == nil,
		"k": validator.ValidateCastling(move.CastlingShort, chess.SideBlack, board) == nil,
		"q": validator.ValidateCastling(move.CastlingLong, chess.SideBlack, board) == nil,
	}

	str := ""
	for castling, isValid := range castlings {
		if isValid {
			str += castling
		}
	}

	if str == "" {
		return nil
	}

	return metric.New("Castlings", str)
}

func EnPassantPosition(board chess.Board) metric.Metric {
	if len(board.MovesHistory()) == 0 {
		return nil
	}

	lastMove := board.MovesHistory()[len(board.MovesHistory())-1]
	normalMove, ok := lastMove.(*move.Normal)
	if !ok || normalMove.PieceNotation != piece.NotationPawn {
		return nil
	}

	if normalMove.To.Rank != normalMove.From.Rank+(piece.PawnRankDirection(!board.Turn())*2) {
		return nil
	}

	passant := position.New(
		normalMove.To.File,
		normalMove.From.Rank+piece.PawnRankDirection(!board.Turn()),
	)

	return metric.New("En passant position", passant)
}
