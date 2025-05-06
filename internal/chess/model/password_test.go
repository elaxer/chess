package model

import "testing"

func TestPassword_Hash(t *testing.T) {
	tests := []struct {
		name    string
		p       Password
		wantErr bool
		isEmpty bool
	}{
		{
			"correct",
			Password("123456"),
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Hash()
			if (err != nil) != tt.wantErr {
				t.Errorf("Password.Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.isEmpty && len(got) <= 0 {
				t.Errorf("Password.Hash() is empty, although must be not")
			}
		})
	}
}
