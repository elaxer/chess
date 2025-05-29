package move_test

import (
	"testing"

	. "github.com/elaxer/chess/pkg/variant/standard/move/move"
)

func TestNewCastling(t *testing.T) {
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
			"short_check",
			args{"0-0+"},
			"0-0+",
			false,
		},
		{
			"short_mate",
			args{"0-0#"},
			"0-0#",
			false,
		},
		{
			"long_check",
			args{"0-0-0+"},
			"0-0-0+",
			false,
		},
		{
			"long_mate",
			args{"0-0-0#"},
			"0-0-0#",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CastlingFromNotation(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCastling() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if got.String() != tt.want {
				t.Errorf("NewCastling() = %v, want %v", got, tt.want)
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
			fields{&Castling{CheckMate: new(CheckMate), CastlingType: CastlingShort}},
			"0-0",
		},
		{
			"castling_long",
			fields{&Castling{CheckMate: new(CheckMate), CastlingType: CastlingLong}},
			"0-0-0",
		},
		{
			"castling_check",
			fields{&Castling{CheckMate: &CheckMate{IsCheck: true, IsMate: false}, CastlingType: CastlingShort}},
			"0-0+",
		},
		{
			"castling_mate",
			fields{&Castling{CheckMate: &CheckMate{IsCheck: false, IsMate: true}, CastlingType: CastlingLong}},
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
