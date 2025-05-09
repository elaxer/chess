package position

import "testing"

func TestRank_Validate(t *testing.T) {
	tests := []struct {
		name    string
		rank    Rank
		wantErr bool
	}{
		{
			"valid",
			Rank4,
			false,
		},
		{
			"less_than_1",
			Rank(0),
			true,
		},
		{
			"bigger_than_max",
			Rank(MaxRank + 1),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rank.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Rank.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
