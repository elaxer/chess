package chess

import "testing"

func TestNewEmptyPosition(t *testing.T) {
	position := NewPositionEmpty()
	if !position.IsEmpty() {
		t.Errorf("NewEmpty() = %v, want Position().IsEmpty() = true", position)
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
			Position{File: FileE, Rank: Rank4},
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
			if got := PositionFromString(tt.args.notation); got != tt.want {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
