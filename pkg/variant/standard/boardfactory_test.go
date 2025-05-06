package standard

import (
	"slices"
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	. "github.com/elaxer/chess/pkg/chess/position"
)

func TestBoardFactory_CreateEmpty(t *testing.T) {
	squares := NewBoardFactory().CreateEmpty().Squares()

	if len(squares) != 64 {
		t.Fatalf("expected 64 squares, got %d", len(squares))
	}
}

// https://www.chess.com/games/view/14842105
func TestBoardFactory_CreateFromMoves(t *testing.T) {
	rawMoves := []RawMove{
		"e4", "c6",
		"d4", "d5",
		"e5", "Bf5",
		"Nc3", "e6",
		"g4", "Bg6",
		"Nge2", "c5",
		"Be3", "Ne7",
		"f4", "h5",
		"f5", "exf5",
		"g5", "Nbc6",
		"Nf4", "a6",
		"Bg2", "cxd4",
		"Bxd4", "Nxd4",
		"Qxd4", "Nc6",
		"Qf2", "Bb4",
		"0-0-0", "Bxc3",
		"bxc3", "Qa5",
		"Rxd5", "Qxc3",
		"Qc5", "Qxc5",
		"Rxc5", "0-0",
		"Bxc6", "bxc6",
		"Rd1", "Rab8",
		"c4", "Rfd8",
		"Rd6", "Kf8",
		"Rcc6", "Rdc8",
		"Kc2", "h4",
		"Rxc8+", "Rxc8",
		"Kc3", "a5",
		"Ra6", "Rb8",
		"Rxa5", "Rb1",
		"c5", "Re1",
		"Ra8+", "Ke7",
		"Ra7+", "Ke8",
		"Nd3", "Re3",
		"Kd2", "Rh3",
		"c6", "Rxh2+",
		"Ke3", "Rc2",
		"e6", "h3",
		"Nb4", "f4+",
		"Kd4", "h2",
		"Ra8+", "Ke7",
		"Rh8", "Rd2+",
		"Kc5", "Be4",
		"c7", "Bb7",
		"Kb6", "Bc8",
		"Rxc8", "h1=Q",
		"Re8+", "Kxe8",
		"c8=Q+", "Ke7",
		"Nc6+", "Qxc6+",
		"Qxc6", "Rd6",
	}
	moves := make([]Move, len(rawMoves))
	for i, move := range rawMoves {
		moves[i] = move
	}

	b, err := NewBoardFactory().CreateFromMoves(moves)
	if err != nil {
		t.Fatalf("failed to create board from moves: %v", err)
	}

	pieces := []struct {
		Position      Position
		PieceNotation PieceNotation
		Side          Side
	}{
		{FromNotation("a2"), NotationPawn, SideWhite},
		{FromNotation("b6"), NotationKing, SideWhite},
		{FromNotation("c6"), NotationQueen, SideWhite},
		{FromNotation("d6"), NotationRook, SideBlack},
		{FromNotation("e6"), NotationPawn, SideWhite},
		{FromNotation("e7"), NotationKing, SideBlack},
		{FromNotation("f4"), NotationPawn, SideBlack},
		{FromNotation("f7"), NotationPawn, SideBlack},
		{FromNotation("g5"), NotationPawn, SideWhite},
		{FromNotation("g7"), NotationPawn, SideBlack},
	}
	positions := make([]Position, len(pieces))
	for i, piece := range pieces {
		positions[i] = piece.Position
	}

	for _, square := range b.Squares() {
		if !slices.Contains(positions, square.Position) && !square.IsEmpty() {
			t.Fatalf("unexpected piece %s on square: %s", square.Piece, square.Position)
		}
	}

	for _, piece := range pieces {
		square := b.Squares().GetByPosition(piece.Position)
		if square.IsEmpty() {
			t.Fatalf("expected piece on square %s, got empty", piece.Position)
		}
		if square.Piece.Notation() != piece.PieceNotation {
			t.Fatalf("expected piece %s, got %s", piece.PieceNotation, square.Piece.Notation())
		}
		if square.Piece.Side() != piece.Side {
			t.Fatalf("expected side %s, got %s", piece.Side, square.Piece.Side())
		}
	}
}
