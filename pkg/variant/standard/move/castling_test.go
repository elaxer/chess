package move

import "testing"

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
			got, err := NewCastling(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewCastling() error = %v, wantErr %v", err, tt.wantErr)
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
		Base *CheckMate
		Type CastlingType
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"castling_short",
			fields{&CheckMate{false, false}, CastlingShort},
			"0-0",
		},
		{
			"castling_long",
			fields{&CheckMate{false, false}, CastlingLong},
			"0-0-0",
		},
		{
			"castling_check",
			fields{&CheckMate{true, false}, CastlingShort},
			"0-0+",
		},
		{
			"castling_mate",
			fields{&CheckMate{false, true}, CastlingLong},
			"0-0-0#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Castling{
				CheckMate:    tt.fields.Base,
				CastlingType: tt.fields.Type,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("Castling.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
