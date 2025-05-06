package move

import (
	"reflect"
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	. "github.com/elaxer/chess/pkg/chess/position"
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
			&Promotion{CheckMate: &CheckMate{}, NewPiece: NotationQueen, To: FromNotation("e8")},
			false,
		},
		{
			"check",
			args{"d1=N+"},
			&Promotion{CheckMate: &CheckMate{IsCheck: true}, NewPiece: NotationKnight, To: FromNotation("d1")},
			false,
		},
		{
			"mate",
			args{"a8=R#"},
			&Promotion{CheckMate: &CheckMate{IsMate: true}, NewPiece: NotationRook, To: FromNotation("a8")},
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
			args{"c7=B"},
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
