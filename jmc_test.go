package jmc

import (
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"testEncode1",
			args{"abc1234"},
			"abc1234",
		},
		{
			"testEncode2",
			args{"abc${enc:def}1234"},
			"abc${enc:wjTxOAveskrorLHjiYyh7g==}1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.s); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"testDecode1",
			args{"abc1234"},
			"abc1234",
			false,
		},
		{
			"testDecode2",
			args{"abc${enc:wjTxOAveskrorLHjiYyh7g==}1234"},
			"abcdef1234",
			false,
		},
		{
			"testDecode3",
			args{"abc${enc:wjTxOAves==}1234"},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
