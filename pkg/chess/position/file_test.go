package position

import (
	"testing"
)

func TestNewFile(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    File
		wantErr bool
	}{
		{
			"a",
			args{"a"},
			FileA,
			false,
		},
		{
			"e",
			args{"e"},
			FileE,
			false,
		},
		{
			"h",
			args{"h"},
			FileH,
			false,
		},
		{
			"x",
			args{"x"},
			FileNull,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileFromString(tt.args.str); got != tt.want {
				t.Errorf("NewFile() got = %v, wantErr %v", got, tt.want)
			}
		})
	}
}

func TestFile_Validate(t *testing.T) {
	tests := []struct {
		name    string
		file    File
		wantErr bool
	}{
		{
			"valid",
			FileE,
			false,
		},
		{
			"null",
			FileNull,
			false,
		},
		{
			"bigger_than_max",
			FileMax + 5,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.file.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("File.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFile_String(t *testing.T) {
	tests := []struct {
		name string
		file File
		want string
	}{
		{
			"a",
			FileA,
			"a",
		},
		{
			"h",
			FileH,
			"h",
		},
		{
			"bigger_than_max",
			FileMax + 5,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.file.String(); got != tt.want {
				t.Errorf("File.String() = %v, wantErr %v", got, tt.want)
			}
		})
	}
}
