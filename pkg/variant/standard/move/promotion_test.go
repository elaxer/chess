package move

import (
	"reflect"
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

func TestNewPromotion(t *testing.T) {
	type args struct {
		notation string
	}
	tests := []struct {
		name    string
		args    args
		want    *Promotion
		wantErr bool
	}{
		{
			"valid",
			args{"e8=Q"},
			&Promotion{CheckMate: &CheckMate{}, NewPiece: NotationQueen, To: position.FromNotation("e8")},
			false,
		},
		{
			"from_file",
			args{"fe8=R"},
			&Promotion{CheckMate: &CheckMate{}, NewPiece: NotationRook, From: position.FromNotation("f"), To: position.FromNotation("e8")},
			false,
		},
		{
			"check",
			args{"d1=N+"},
			&Promotion{CheckMate: &CheckMate{IsCheck: true}, NewPiece: NotationKnight, To: position.FromNotation("d1")},
			false,
		},
		{
			"mate",
			args{"a8=R#"},
			&Promotion{CheckMate: &CheckMate{IsMate: true}, NewPiece: NotationRook, To: position.FromNotation("a8")},
			false,
		},
		{
			"invalid_piece",
			args{"c1=K"},
			nil,
			true,
		},
		{
			"invalid_file",
			args{"i8=B"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPromotion(tt.args.notation)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewPromotion() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPromotion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromotion_Notation(t *testing.T) {
	type fields struct {
		CheckMate *CheckMate
		From      position.Position
		To        position.Position
		NewPiece  PieceNotation
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"valid",
			fields{CheckMate: &CheckMate{}, To: position.FromNotation("a1"), NewPiece: NotationRook},
			"a1=R",
		},
		{
			"from_file",
			fields{CheckMate: &CheckMate{}, From: position.FromNotation("f"), To: position.FromNotation("e8"), NewPiece: NotationRook},
			"fe8=R",
		},
		{
			"full_from",
			fields{CheckMate: &CheckMate{}, From: position.FromNotation("b2"), To: position.FromNotation("b1"), NewPiece: NotationKnight},
			"b2b1=N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Promotion{
				CheckMate: tt.fields.CheckMate,
				From:      tt.fields.From,
				To:        tt.fields.To,
				NewPiece:  tt.fields.NewPiece,
			}
			if got := m.Notation(); got != tt.want {
				t.Errorf("Promotion.Notation() = %v, want %v", got, tt.want)
			}
		})
	}
}
