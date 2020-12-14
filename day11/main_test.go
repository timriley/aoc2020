package main

import (
	"reflect"
	"testing"
)

func Test_room_neighbours(t *testing.T) {
	type fields struct {
		cells [][]string
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "centered",
			fields: fields{
				cells: [][]string{
					{"L", ".", "L"},
					{".", ".", "L"},
					{"L", ".", "."},
				},
			},
			args: args{x: 1, y: 1},
			want: []string{"L", ".", "L", ".", "L", "L", ".", "."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &room{
				cells: tt.fields.cells,
			}
			if got := r.neighbours(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neighbours() = %v, want %v", got, tt.want)
			}
		})
	}
}
