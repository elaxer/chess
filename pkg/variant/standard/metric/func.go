package metric

import (
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/metric"
	"github.com/elaxer/chess/pkg/variant/standard/move/move"
	"github.com/elaxer/chess/pkg/variant/standard/move/resolver"
	"github.com/elaxer/chess/pkg/variant/standard/move/validator"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

type Castlings = map[string]map[chess.Side]map[move.CastlingType]bool

var AllFuncs = []metric.MetricFunc{
	CastlingAbility,
	EnPassantTargetSquare,
	HalfmoveClock,
}

func CastlingAbility(board chess.Board) metric.Metric {
	callback := func(side chess.Side, board chess.Board, validateObstacle bool) map[move.CastlingType]bool {
		return map[move.CastlingType]bool{
			move.CastlingShort: validator.ValidateCastling(move.CastlingShort, side, board, validateObstacle) == nil,
			move.CastlingLong:  validator.ValidateCastling(move.CastlingLong, side, board, validateObstacle) == nil,
		}
	}

	castlings := Castlings{
		"theoretical": {
			chess.SideWhite: callback(chess.SideWhite, board, false),
			chess.SideBlack: callback(chess.SideBlack, board, false),
		},
		"practical": {
			chess.SideWhite: callback(chess.SideWhite, board, true),
			chess.SideBlack: callback(chess.SideBlack, board, true),
		},
	}

	return metric.New("Castling ability", castlings)
}

// todo переделать
func EnPassantTargetSquare(board chess.Board) metric.Metric {
	if len(board.MovesHistory()) == 0 {
		return nil
	}

	lastMove := board.MovesHistory()[len(board.MovesHistory())-1]
	normalMove, ok := lastMove.(*move.Normal)
	if !ok || normalMove.PieceNotation != piece.NotationPawn {
		return nil
	}

	from, err := resolver.ResolveFrom(normalMove, board, board.Turn())
	if err != nil {
		return nil
	}

	if normalMove.To.Rank != from.Rank+(piece.PawnRankDirection(!board.Turn())*2) {
		return nil
	}

	passant := position.New(
		normalMove.To.File,
		from.Rank+piece.PawnRankDirection(!board.Turn()),
	)

	return metric.New("En passant target square", passant)
}

func HalfmoveClock(board chess.Board) metric.Metric {
	clock := 0
	for _, m := range board.MovesHistory() {
		normalMove, ok := m.(*move.Normal)
		if !ok || normalMove.PieceNotation == piece.NotationPawn || normalMove.IsCapture {
			clock = 0

			continue
		}

		clock++
	}

	return metric.New("Halfmove clock", clock)
}
