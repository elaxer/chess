package handler

import (
	"net/url"
	"testing"
)

func TestParamQueryInt(t *testing.T) {
	type args struct {
		query     url.Values
		key       string
		byDefault int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"correct",
			args{url.Values{"q": []string{"5"}}, "q", 10},
			5,
		},
		{
			"empty",
			args{url.Values{}, "q", 10},
			10,
		},
		{
			"non integer",
			args{url.Values{"q": []string{"q"}}, "q", 10},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamQueryInt(tt.args.query, tt.args.key, tt.args.byDefault); got != tt.want {
				t.Errorf("ParamQueryInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
