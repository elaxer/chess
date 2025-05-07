package move

import (
	"testing"

	. "github.com/elaxer/chess/pkg/chess"
	"github.com/elaxer/chess/pkg/chess/position"
)

func TestNewNormal(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"pawn",
			args{"e4"},
			"e4",
			false,
		},
		{
			"rook",
			args{"Rd8"},
			"Rd8",
			false,
		},
		{
			"bishop",
			args{"Ba1"},
			"Ba1",
			false,
		},
		{
			"knight",
			args{"Nc3"},
			"Nc3",
			false,
		},
		{
			"queen",
			args{"Qc6"},
			"Qc6",
			false,
		},
		{
			"king",
			args{"Kb7"},
			"Kb7",
			false,
		},
		{
			"check",
			args{"Rg1+"},
			"Rg1+",
			false,
		},
		{
			"mate",
			args{"Be8#"},
			"Be8#",
			false,
		},
		{
			"capture",
			args{"Rxh7"},
			"Rxh7",
			false,
		},
		{
			"capture_check",
			args{"xb5+"},
			"xb5+",
			false,
		},
		{
			"capture_mate",
			args{"Qxh8#"},
			"Qxh8#",
			false,
		},
		{
			"error",
			args{"Ik9"},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNormal(tt.args.str)

			if (err != nil) != tt.wantErr {
				t.Fatalf("NewNormal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			if gotStr := got.String(); gotStr != tt.want {
				t.Errorf("NewNormal() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func TestNormal_String(t *testing.T) {
	type fields struct {
		Position  position.Position
		Piece     PieceNotation
		IsCheck   bool
		IsMate    bool
		IsCapture bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"normal",
			fields{position.FromNotation("a8"), NotationQueen, false, false, false},
			"Qa8",
		},
		{
			"normal_pawn",
			fields{position.FromNotation("e4"), NotationPawn, false, false, false},
			"e4",
		},
		{
			"check",
			fields{position.FromNotation("a1"), NotationRook, true, false, false},
			"Ra1+",
		},
		{
			"mate",
			fields{position.FromNotation("a1"), NotationBishop, false, true, false},
			"Ba1#",
		},
		{
			"check_capture",
			fields{position.FromNotation("a1"), NotationKnight, true, false, true},
			"Nxa1+",
		},
		{
			"mate_capture",
			fields{position.FromNotation("c5"), NotationPawn, false, true, true},
			"xc5#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			move := Normal{
				CheckMate:     &CheckMate{tt.fields.IsCheck, tt.fields.IsMate},
				To:            tt.fields.Position,
				PieceNotation: tt.fields.Piece,
				IsCapture:     tt.fields.IsCapture,
			}
			if got := move.String(); got != tt.want {
				t.Errorf("Normal.String() = %v, wantErr %v", got, tt.want)
			}
		})
	}
}
