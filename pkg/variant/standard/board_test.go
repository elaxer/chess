package standard

import (
	"testing"

	"github.com/elaxer/chess/pkg/chess"
	. "github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
)

func TestBoard_isCheck(t *testing.T) {
	type fields struct {
		Turn    chess.Side
		Squares []*chess.Square
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"check",
			fields{
				chess.SideWhite,
				[]*chess.Square{
					{Position: FromNotation("a1"), Piece: piece.NewKing(chess.SideWhite)},
					{Position: FromNotation("h8"), Piece: piece.NewKing(chess.SideBlack)},
					{Position: FromNotation("a8"), Piece: piece.NewRook(chess.SideBlack)},
				},
			},
			true,
		},
		{
			"check_bishop",
			fields{
				chess.SideBlack,
				[]*chess.Square{
					{Position: FromNotation("e1"), Piece: piece.NewKing(chess.SideBlack)},
					{Position: FromNotation("h8"), Piece: piece.NewKing(chess.SideWhite)},
					{Position: FromNotation("b4"), Piece: piece.NewBishop(chess.SideWhite)},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ((&boardFactory{}).CreateEmpty()).(*board)
			if tt.fields.Turn == chess.SideBlack {
				b.NextTurn()
			}
			for _, square := range tt.fields.Squares {
				b.squares.AddPiece(square.Piece, square.Position)
			}
			if got := b.isCheck(b.turn); got != tt.want {
				t.Errorf("Board.isCheck() = %v, wantErr %v", got, tt.want)
			}
		})
	}
}

// todo : debug testf
func TestBoard_isCheckOpponent(t *testing.T) {
	type fields struct {
		Turn    chess.Side
		Squares []*chess.Square
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"check",
			fields{
				chess.SideBlack,
				[]*chess.Square{
					{Position: FromNotation("a1"), Piece: piece.NewKing(chess.SideWhite)},
					{Position: FromNotation("h8"), Piece: piece.NewKing(chess.SideBlack)},
					{Position: FromNotation("a8"), Piece: piece.NewRook(chess.SideBlack)},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ((&boardFactory{}).CreateEmpty()).(*board)
			if tt.fields.Turn == chess.SideBlack {
				b.NextTurn()
			}
			for _, square := range tt.fields.Squares {
				b.squares.AddPiece(square.Piece, square.Position)
			}
			if got := b.isCheck(!b.turn); got != tt.want {
				t.Errorf("Board.isCheck() = %v, wantErr %v", got, tt.want)
			}
		})
	}
}

func TestBoard_isMate(t *testing.T) {
	type fields struct {
		Turn    chess.Side
		Squares []*chess.Square
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"mate",
			fields{
				chess.SideWhite,
				[]*chess.Square{
					{Position: FromNotation("a1"), Piece: piece.NewKing(chess.SideWhite)},
					{Position: FromNotation("h8"), Piece: piece.NewKing(chess.SideBlack)},
					{Position: FromNotation("a8"), Piece: piece.NewRook(chess.SideBlack)},
					{Position: FromNotation("b8"), Piece: piece.NewRook(chess.SideBlack)},
				},
			},
			true,
		},
		{
			"no_mate_because_king_can_capture_rook",
			fields{
				chess.SideWhite,
				[]*chess.Square{
					{Position: FromNotation("a1"), Piece: piece.NewKing(chess.SideWhite)},
					{Position: FromNotation("h8"), Piece: piece.NewKing(chess.SideBlack)},
					{Position: FromNotation("a2"), Piece: piece.NewRook(chess.SideBlack)},
					{Position: FromNotation("b8"), Piece: piece.NewRook(chess.SideBlack)},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ((&boardFactory{}).CreateEmpty()).(*board)
			if tt.fields.Turn == chess.SideBlack {
				b.NextTurn()
			}
			for _, square := range tt.fields.Squares {
				b.Squares().AddPiece(square.Piece, square.Position)
			}

			if got := b.State().IsMate(); got != tt.want {
				t.Errorf("Board.isMate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_isStalemate(t *testing.T) {
	type fields struct {
		Turn    chess.Side
		Squares []*chess.Square
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"stalemate",
			fields{
				chess.SideWhite,
				[]*chess.Square{
					{Position: FromNotation("a8"), Piece: piece.NewKing(chess.SideWhite)},
					{Position: FromNotation("h1"), Piece: piece.NewKing(chess.SideBlack)},
					{Position: FromNotation("c7"), Piece: piece.NewQueen(chess.SideBlack)},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ((&boardFactory{}).CreateEmpty()).(*board)
			if tt.fields.Turn == chess.SideBlack {
				b.NextTurn()
			}
			for _, square := range tt.fields.Squares {
				b.Squares().AddPiece(square.Piece, square.Position)
			}

			if got := b.State().IsStalemate(); got != tt.want {
				t.Errorf("Board.isStalemate() = %v, wantErr %v", got, tt.want)
			}
		})
	}
}
