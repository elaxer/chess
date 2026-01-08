package chess

import (
	"testing"
)

func TestNewPositionEmpty(t *testing.T) {
	position := NewPositionEmpty()
	if !position.IsEmpty() {
		t.Errorf("NewPositionEmpty() = %v, want Position().IsEmpty() = true", position)
	}
}

func TestPositionFromString(t *testing.T) {
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
			Position{File: FileE, Rank: Rank4},
		},
		{
			"full_two_digits_rank",
			args{"e14"},
			Position{File: FileE, Rank: Rank14},
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
		{
			"two_digits_rank",
			args{"12"},
			Position{Rank: Rank12},
		},
		{
			"empty",
			args{""},
			Position{},
		},
		{
			"outranged_values",
			args{"x22"},
			Position{},
		},
		{
			"outranged_file",
			args{"z"},
			Position{},
		},
		{
			"outranged_rank",
			args{"32"},
			Position{},
		},
		{
			"wrong_format",
			args{"foo223"},
			Position{},
		},
		{
			"utf8",
			args{"ж3"},
			Position{},
		},
		{
			"emoji",
			args{"🇰🇿8"},
			Position{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PositionFromString(tt.args.notation); got != tt.want {
				t.Errorf("PositionFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
