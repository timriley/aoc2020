package main

import (
	"reflect"
	"testing"
)

func Test_joltageDifferences(t *testing.T) {
	type args struct {
		adapters []int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			"Simple case",
			args{
				adapters: []int{
					16,
					10,
					15,
					5,
					1,
					11,
					7,
					19,
					6,
					12,
					4,
				},
			},
			map[int]int{1: 7, 3: 5},
		},
		{
			"Complex case",
			args{
				adapters: []int{
					28,
					33,
					18,
					42,
					31,
					14,
					46,
					20,
					48,
					47,
					24,
					23,
					49,
					45,
					19,
					38,
					39,
					11,
					1,
					32,
					25,
					35,
					8,
					17,
					7,
					9,
					4,
					2,
					34,
					10,
					3,
				},
			},
			map[int]int{1: 22, 3: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joltageDifferences(tt.args.adapters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("joltageDifferences() = %v, want %v", got, tt.want)
			}
		})
	}
}
