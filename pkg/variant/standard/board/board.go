package board

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/metric"
	"github.com/elaxer/chess/pkg/variant/standard/move/mover"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state/rule"
)

var edgePosition = position.New(position.FileH, position.Rank8)

type board struct {
	turn           chess.Side
	squares        *chess.Squares
	movesHistory   []chess.MoveResult
	capturedPieces []chess.Piece

	stateRules []rule.Rule
}

func (b *board) Squares() *chess.Squares {
	return b.squares
}

func (b *board) Turn() chess.Side {
	return b.turn
}

func (b *board) State(side chess.Side) chess.State {
	for _, rule := range b.stateRules {
		if state := rule(b, side); state != nil {
			return state
		}
	}

	return chess.StateClear
}

func (b *board) MovesHistory() []chess.MoveResult {
	return b.movesHistory
}

func (b *board) Moves(side chess.Side) position.Set {
	moves := mapset.NewSetWithSize[position.Position](32)
	for _, piece := range b.squares.GetAllPieces(side) {
		moves = moves.Union(b.LegalMoves(piece))
	}

	return moves
}

func (b *board) LegalMoves(p chess.Piece) position.Set {
	from := b.squares.GetByPiece(p)
	if from.IsEmpty() {
		return nil
	}

	pseudoMoves := p.PseudoMoves(from, b.squares)

	if p.Side() != b.Turn() {
		return pseudoMoves
	}

	legalMoves := mapset.NewSetWithSize[position.Position](pseudoMoves.Cardinality())
	for to := range pseudoMoves.Iter() {
		b.squares.MovePieceTemporarily(from, to, func() {
			_, kingPosition := b.squares.FindPiece(piece.NotationKing, b.turn)
			if !b.Moves(!b.turn).ContainsOne(kingPosition) {
				legalMoves.Add(to)
			}
		})
	}

	return legalMoves
}

func (b *board) MakeMove(move chess.Move) (chess.MoveResult, error) {
	moveResult, err := mover.MakeMove(move, b)
	if err != nil {
		return nil, err
	}

	b.movesHistory = append(b.movesHistory, moveResult)
	b.turn = !b.turn

	return moveResult, nil
}

func (b *board) UndoLastMove() (chess.MoveResult, error) {
	return nil, nil
}

func (b *board) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"squares":         b.squares,
		"state":           b.State(b.turn),
		"captured_pieces": b.capturedPieces,
		"castlings":       metric.CastlingAbility(b).Value().(metric.Castlings)["practical"][b.turn],
	})
}
