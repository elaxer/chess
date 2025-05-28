package standard_test

import (
	"testing"

	"github.com/elaxer/chess/pkg/chess"
	. "github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state"
	"github.com/elaxer/chess/pkg/variant/standardtest"
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
			fields{standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
				{Piece: piece.NewKing(chess.SideWhite), Position: FromNotation("a1")},
				{Piece: piece.NewKing(chess.SideBlack), Position: FromNotation("h8")},
				{Piece: piece.NewRook(chess.SideBlack), Position: FromNotation("a8")},
			})},
			state.Check,
		},
		{
			"check_bishop",
			fields{standardtest.NewEmpty(chess.SideBlack, []standardtest.Placement{
				{Piece: piece.NewKing(chess.SideBlack), Position: FromNotation("e1")},
				{Piece: piece.NewKing(chess.SideWhite), Position: FromNotation("h8")},
				{Piece: piece.NewBishop(chess.SideWhite), Position: FromNotation("b4")},
			})},
			state.Check,
		},

		{
			"mate",
			fields{standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
				{Piece: piece.NewKing(chess.SideWhite), Position: FromNotation("a1")},
				{Piece: piece.NewKing(chess.SideBlack), Position: FromNotation("h8")},
				{Piece: piece.NewRook(chess.SideBlack), Position: FromNotation("a8")},
				{Piece: piece.NewRook(chess.SideBlack), Position: FromNotation("b8")},
			})},
			state.Mate,
		},
		{
			// no mate because the black king can capture the threatening rook
			"no_mate",
			fields{standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
				{Piece: piece.NewKing(chess.SideWhite), Position: FromNotation("a1")},
				{Piece: piece.NewKing(chess.SideBlack), Position: FromNotation("h8")},
				{Piece: piece.NewRook(chess.SideBlack), Position: FromNotation("a2")},
				{Piece: piece.NewRook(chess.SideBlack), Position: FromNotation("b8")},
			})},
			state.Check,
		},

		{
			"stalemate",
			fields{standardtest.NewEmpty(chess.SideWhite, []standardtest.Placement{
				{Piece: piece.NewKing(chess.SideWhite), Position: FromNotation("a8")},
				{Piece: piece.NewKing(chess.SideBlack), Position: FromNotation("b6")},
				{Piece: piece.NewQueen(chess.SideBlack), Position: FromNotation("c7")},
			})},
			state.Stalemate,
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
