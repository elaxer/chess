package standard_test

import (
	"testing"

	"github.com/elaxer/chess/pkg/chess"
	. "github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standarttest"
)

func Test_standard_State(t *testing.T) {
	type fields struct {
		board chess.Board
	}
	tests := []struct {
		name   string
		fields fields
		want   chess.State
	}{
		{
			"check",
			fields{standarttest.NewEmpty(chess.SideWhite, []standarttest.Placement{
				{Piece: piece.NewKing(chess.SideWhite), Position: FromNotation("a1")},
				{Piece: piece.NewKing(chess.SideBlack), Position: FromNotation("h8")},
				{Piece: piece.NewRook(chess.SideBlack), Position: FromNotation("a8")},
			})},
			chess.StateCheck,
		},
		{
			"check_bishop",
			fields{standarttest.NewEmpty(chess.SideBlack, []standarttest.Placement{
				{Piece: piece.NewKing(chess.SideBlack), Position: FromNotation("e1")},
				{Piece: piece.NewKing(chess.SideWhite), Position: FromNotation("h8")},
				{Piece: piece.NewBishop(chess.SideWhite), Position: FromNotation("b4")},
			})},
			chess.StateCheck,
		},

		{
			"mate",
			fields{standarttest.NewEmpty(chess.SideWhite, []standarttest.Placement{
				{Piece: piece.NewKing(chess.SideWhite), Position: FromNotation("a1")},
				{Piece: piece.NewKing(chess.SideBlack), Position: FromNotation("h8")},
				{Piece: piece.NewRook(chess.SideBlack), Position: FromNotation("a8")},
				{Piece: piece.NewRook(chess.SideBlack), Position: FromNotation("b8")},
			})},
			chess.StateMate,
		},
		{
			// нет мата, потому что черный король может забрать угрожающую ладью
			"no_mate",
			fields{standarttest.NewEmpty(chess.SideWhite, []standarttest.Placement{
				{Piece: piece.NewKing(chess.SideWhite), Position: FromNotation("a1")},
				{Piece: piece.NewKing(chess.SideBlack), Position: FromNotation("h8")},
				{Piece: piece.NewRook(chess.SideBlack), Position: FromNotation("a2")},
				{Piece: piece.NewRook(chess.SideBlack), Position: FromNotation("b8")},
			})},
			chess.StateCheck,
		},

		{
			"stalemate",
			fields{standarttest.NewEmpty(chess.SideWhite, []standarttest.Placement{
				{Piece: piece.NewKing(chess.SideWhite), Position: FromNotation("a8")},
				{Piece: piece.NewKing(chess.SideBlack), Position: FromNotation("h1")},
				{Piece: piece.NewQueen(chess.SideBlack), Position: FromNotation("c7")},
			})},
			chess.StateStalemate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.board.State(tt.fields.board.Turn()); got != tt.want {
				t.Errorf("standard.State() = %v, want %v", got, tt.want)
			}
		})
	}
}
