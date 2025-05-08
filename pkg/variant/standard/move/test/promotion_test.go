package move_test

import (
	"reflect"
	"testing"

	. "github.com/elaxer/chess/pkg/variant/standard/move"

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
			&Promotion{
				Normal:   &Normal{CheckMate: new(CheckMate), To: position.FromNotation("e8")},
				NewPiece: NotationQueen,
			},
			false,
		},
		{
			"from_file",
			args{"fe8=R"},
			&Promotion{
				Normal:   &Normal{CheckMate: new(CheckMate), From: position.FromNotation("f"), To: position.FromNotation("e8")},
				NewPiece: NotationRook,
			},
			false,
		},
		{
			"check",
			args{"d1=N+"},
			&Promotion{
				Normal:   &Normal{CheckMate: &CheckMate{IsCheck: true}, To: position.FromNotation("d1")},
				NewPiece: NotationKnight,
			},
			false,
		},
		{
			"mate",
			args{"a8=R#"},
			&Promotion{
				Normal:   &Normal{CheckMate: &CheckMate{IsMate: true}, To: position.FromNotation("a8")},
				NewPiece: NotationRook,
			},
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
				t.Errorf("NewPromotion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			got.Normal.PieceNotation = NotationPawn

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPromotion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromotion_Notation(t *testing.T) {
	type fields struct {
		promotion *Promotion
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"valid",
			fields{&Promotion{
				Normal:   &Normal{CheckMate: new(CheckMate), To: position.FromNotation("a1")},
				NewPiece: NotationRook,
			}},
			"a1=R",
		},
		{
			"from_file",
			fields{
				&Promotion{
					Normal:   &Normal{CheckMate: new(CheckMate), From: position.FromNotation("f"), To: position.FromNotation("e8")},
					NewPiece: NotationRook,
				},
			},
			"fe8=R",
		},
		{
			"full_from",
			fields{&Promotion{
				Normal:   &Normal{CheckMate: new(CheckMate), From: position.FromNotation("b2"), To: position.FromNotation("b1")},
				NewPiece: NotationKnight,
			}},
			"b2b1=N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.promotion.Notation(); got != tt.want {
				t.Errorf("Promotion.Notation() = %v, want %v", got, tt.want)
			}
		})
	}
}
