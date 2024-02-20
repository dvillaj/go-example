package main

import (
	"testing"
)

func Test_randInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "CheckValue In Range",
			args: args{
				min: 1,
				max: 2,
			},
			want: 1,
		},
		{name: "CheckValue In Range Larger",
			args: args{
				min: 99,
				max: 100,
			},
			want: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randInt(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("randInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printMessage(t *testing.T) {
	type args struct {
		msgNum int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printMessage(tt.args.msgNum)
		})
	}
}
