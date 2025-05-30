package move

import (
	"testing"

	"github.com/elaxer/chess/pkg/chess/position"
	"github.com/elaxer/chess/pkg/variant/standard/piece"
	"github.com/elaxer/chess/pkg/variant/standard/state/state"
)

func TestNormalFromNotation(t *testing.T) {
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
			got, err := NormalFromNotation(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalFromNotation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if gotStr := got.String(); gotStr != tt.want {
				t.Errorf("NormalFromNotation().String() = %v, want %v", gotStr, tt.want)
			}
		})
	}
}

func TestNormal_String(t *testing.T) {
	tests := []struct {
		name string
		move *Normal
		want string
	}{
		{
			"normal",
			&Normal{
				To:            position.FromNotation("a8"),
				PieceNotation: piece.NotationQueen,
				IsCapture:     false,
			},
			"Qa8",
		},
		{
			"pawn",
			&Normal{
				To:            position.FromNotation("e4"),
				PieceNotation: piece.NotationPawn,
				IsCapture:     false,
			},
			"e4",
		},
		{
			"check",
			&Normal{
				abstract:      abstract{NewBoardState: state.Check},
				To:            position.FromNotation("a1"),
				PieceNotation: piece.NotationRook,
				IsCapture:     false,
			},
			"Ra1+",
		},
		{
			"mate",
			&Normal{
				abstract:      abstract{NewBoardState: state.Mate},
				To:            position.FromNotation("a1"),
				PieceNotation: piece.NotationBishop,
				IsCapture:     false,
			},
			"Ba1#",
		},
		{
			"check_with_capture",
			&Normal{
				abstract:      abstract{NewBoardState: state.Check},
				To:            position.FromNotation("a1"),
				PieceNotation: piece.NotationKnight,
				IsCapture:     true,
			},
			"Nxa1+",
		},
		{
			"mate_with_capture",
			&Normal{
				abstract:      abstract{NewBoardState: state.Mate},
				To:            position.FromNotation("c5"),
				PieceNotation: piece.NotationPawn,
				IsCapture:     true,
			},
			"xc5#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.move.String(); got != tt.want {
				t.Errorf("Normal.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
