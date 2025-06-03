package position

import (
	"testing"
)

func TestNewNull(t *testing.T) {
	position := NewEmpty()
	if !position.IsEmpty() {
		t.Errorf("NewNull() = %v, want Position().IsNull() = true", position)
	}
}

func TestFromString(t *testing.T) {
	type args struct {
		notation string
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{
			"full",
			args{"e4"},
			Position{FileE, Rank4},
		},
		{
			"file",
			args{"c"},
			Position{File: FileC},
		},
		{
			"rank",
			args{"3"},
			Position{Rank: Rank3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromString(tt.args.notation); got != tt.want {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			"full",
			fields{FileA, Rank1},
			args{`{"file": 1, "rank": 1}`},
			false,
		},
		{
			"only_file",
			fields{File: FileE},
			args{`{"file": 5}`},
			false,
		},
		{
			"only_rank",
			fields{Rank: Rank9},
			args{`{"rank": 9}`},
			false,
		},
		{
			"invalid",
			fields{},
			args{`{"file": -3, "rank": 100}`},
			false,
		},
		{
			"invalid_file",
			fields{Rank: Rank2},
			args{`{"file": 18, "rank": 2}`},
			false,
		},
		{
			"invalid_rank",
			fields{File: FileC},
			args{`{"file": 3, "rank": 0}`},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			position := NewEmpty()
			if err := position.UnmarshalJSON([]byte(tt.args.data)); (err != nil) != tt.wantErr {
				t.Errorf("Position.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if position.File != tt.fields.File || position.Rank != tt.fields.Rank {
				t.Errorf("Position is %v, expected %v", position, Position{tt.fields.File, tt.fields.Rank})
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
			fields{FileE, Rank4},
			"e4",
		},
		{
			"nulls",
			fields{FileNull, RankNull},
			"",
		},
		{
			"bigger_than_max_values",
			fields{File(FileMax + 1), Rank(RankMax + 1)},
			"",
		},
		{
			"null_file",
			fields{FileNull, Rank4},
			"4",
		},
		{
			"null_rank",
			fields{FileD, RankNull},
			"d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.fields.File, tt.fields.Rank).String(); got != tt.want {
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
			fields{24, 89},
			true,
		},
		{
			"invalid_rank",
			fields{FileA, 17},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := New(tt.fields.File, tt.fields.Rank).Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Position.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPosition_IsInRange(t *testing.T) {
	type fields struct {
		File File
		Rank Rank
	}
	type args struct {
		position Position
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"invalid",
			fields{FileI, Rank2},
			args{New(FileH, Rank8)},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Position{
				File: tt.fields.File,
				Rank: tt.fields.Rank,
			}
			if got := p.IsBoundaries(tt.args.position); got != tt.want {
				t.Errorf("Position.IsInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
