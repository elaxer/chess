package position

import (
	"testing"
)

func TestPosition_UnmarshalJSON(t *testing.T) {
	type fields struct {
		File File
		Rank Rank
	}
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"valid",
			fields{FileA, Rank1},
			args{`{"file": 1, "rank": 1}`},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			position := Position{}

			if err := position.UnmarshalJSON([]byte(tt.args.data)); (err != nil) != tt.wantErr {
				t.Errorf("Position.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if position.File != tt.fields.File {
				t.Errorf("Position.UnmarshalJSON() got = %v, wantErr %v", position.File, tt.fields.File)
				return
			}
			if position.Rank != tt.fields.Rank {
				t.Errorf("Position.UnmarshalJSON() c.Rank = %v, wantErr %v", position.Rank, tt.fields.Rank)
			}
		})
	}
}

func TestPosition_String(t *testing.T) {
	type fields struct {
		File File
		Rank Rank
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"valid",
			fields{File(5), Rank(4)},
			"e4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			position := Position{
				File: tt.fields.File,
				Rank: tt.fields.Rank,
			}
			if got := position.String(); got != tt.want {
				t.Errorf("Position.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_Validate(t *testing.T) {
	type fields struct {
		File File
		Rank Rank
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"valid",
			fields{FileA, Rank1},
			false,
		},
		{
			"invalid",
			fields{0, Rank8},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				File: tt.fields.File,
				Rank: tt.fields.Rank,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Position.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
