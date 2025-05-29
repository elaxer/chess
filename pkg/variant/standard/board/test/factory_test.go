package board_test

import (
	"testing"

	"github.com/elaxer/chess/pkg/variant/standard/board"
	"github.com/elaxer/chess/pkg/variant/standard/codec/fen"
	"github.com/elaxer/chess/pkg/variant/standardtest"
)

// https://www.chess.com/games/view/14842105
func TestFactory_CreateFromMoves(t *testing.T) {
	moves := []string{
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

	b, err := board.NewFactory().CreateFromMoves(standardtest.NotationsToMoves(moves))
	if err != nil {
		t.Fatalf("failed to create board from moves: %v", err)
	}

	fen := fen.Encode(b)
	const expectedFEN = "8/4kpp1/1KQrP3/6P1/5p2/8/P7/8 w - - 1 51"
	if fen != expectedFEN {
		t.Errorf("Expected position \"%s\", got - %s", expectedFEN, fen)
	}
}
