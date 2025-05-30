package move

import (
	"testing"

	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"

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
			"promotion",
			args{"e8=Q"},
			&Promotion{
				Normal:           &Normal{To: position.FromNotation("e8")},
				NewPieceNotation: piece.NotationQueen,
			},
			false,
		},
		{
			"from_file",
			args{"fe8=R"},
			&Promotion{
				Normal:           &Normal{From: position.FromNotation("f"), To: position.FromNotation("e8")},
				NewPieceNotation: piece.NotationRook,
			},
			false,
		},
		{
			"check",
			args{"d1=N+"},
			&Promotion{
				Normal:           &Normal{abstract: abstract{NewBoardState: state.Check}, To: position.FromNotation("d1")},
				NewPieceNotation: piece.NotationKnight,
			},
			false,
		},
		{
			"mate",
			args{"a8=R#"},
			&Promotion{
				Normal:           &Normal{abstract: abstract{NewBoardState: state.Mate}, To: position.FromNotation("a8")},
				NewPieceNotation: piece.NotationRook,
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
			args{"w8=B"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PromotionFromNotation(tt.args.notation)
			if (err != nil) != tt.wantErr {
				t.Errorf("PromotionFromNotation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			got.Normal.PieceNotation = piece.NotationPawn

			if got.String() != tt.want.String() {
				t.Errorf("PromotionFromNotation().String() = %v, want %v", got, tt.want)
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
			"promotion",
			fields{&Promotion{
				Normal:           &Normal{To: position.FromNotation("a1")},
				NewPieceNotation: piece.NotationRook,
			}},
			"a1=R",
		},
		{
			"from_file",
			fields{
				&Promotion{
					Normal:           &Normal{From: position.FromNotation("f"), To: position.FromNotation("e8")},
					NewPieceNotation: piece.NotationRook,
				},
			},
			"fe8=R",
		},
		{
			"full_from",
			fields{&Promotion{
				Normal:           &Normal{From: position.FromNotation("b2"), To: position.FromNotation("b1")},
				NewPieceNotation: piece.NotationKnight,
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
