package main

import (
	"reflect"
	"testing"
)

func Test_processGrid(t *testing.T) {
	type args struct {
		grid [][]bool
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{
			name: "tests next state 1",
			args: args{
				grid: [][]bool{

					// {true, true, false, true, true},
					// {false, false, false, true, false},
					// {true, false, true, true, false},
					// {true, false, false, false, false},
					// {true, true, false, true, false},
					{false, false, true},
					{false, true, true},
					{false, false, false},
				},
			},
			want: [][]bool{
				{false, true, true},
				{false, true, true},
				{false, false, false},
				// 	{false, false, false, true, true},
				// 	{false, false, t, true, t},

				// 	{true, false, true, true, false},
				// 	{true, false, false, false, false},
				// 	{true, true, false, true, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processGrid(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
