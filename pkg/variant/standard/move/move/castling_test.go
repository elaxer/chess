package move

import (
	"testing"

	"github.com/elaxer/chess/pkg/variant/standard/state/state"
)

func TestCastlingFromNotation(t *testing.T) {
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
			"short",
			args{"0-0"},
			"0-0",
			false,
		},
		{
			"long",
			args{"0-0-0"},
			"0-0-0",
			false,
		},
		{
			"short_with_check",
			args{"0-0+"},
			"0-0+",
			false,
		},
		{
			"short_with_mate",
			args{"0-0#"},
			"0-0#",
			false,
		},
		{
			"long_with_check",
			args{"0-0-0+"},
			"0-0-0+",
			false,
		},
		{
			"long_with_mate",
			args{"0-0-0#"},
			"0-0-0#",
			false,
		},
		{
			"O character",
			args{"O-O"},
			"0-0",
			false,
		},
		{
			"All characters",
			args{"O-o-0+"},
			"0-0-0+",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CastlingFromNotation(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("CastlingFromNotation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("CastlingFromNotation().String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCastling_String(t *testing.T) {
	type fields struct {
		move *Castling
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"castling_short",
			fields{NewCastling(CastlingShort)},
			"0-0",
		},
		{
			"castling_long",
			fields{NewCastling(CastlingLong)},
			"0-0-0",
		},
		{
			"castling_with_check",
			fields{&Castling{abstract: abstract{NewBoardState: state.Check}, CastlingType: CastlingShort}},
			"0-0+",
		},
		{
			"castling_with_mate",
			fields{&Castling{abstract: abstract{NewBoardState: state.Mate}, CastlingType: CastlingLong}},
			"0-0-0#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.move.String(); got != tt.want {
				t.Errorf("Castling.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
