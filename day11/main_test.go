package main

import (
	"fmt"
	"reflect"
	"strings"
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

func Test_room_iterate(t *testing.T) {
	stages := [][][]string{
		[][]string{
			[]string{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
			[]string{"L", "L", "L", "L", "L", "L", "L", ".", "L", "L"},
			[]string{"L", ".", "L", ".", "L", ".", ".", "L", ".", "."},
			[]string{"L", "L", "L", "L", ".", "L", "L", ".", "L", "L"},
			[]string{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
			[]string{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
			[]string{".", ".", "L", ".", "L", ".", ".", ".", ".", "."},
			[]string{"L", "L", "L", "L", "L", "L", "L", "L", "L", "L"},
			[]string{"L", ".", "L", "L", "L", "L", "L", "L", ".", "L"},
			[]string{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
		},
		[][]string{
			[]string{"#", ".", "#", "#", ".", "#", "#", ".", "#", "#"},
			[]string{"#", "#", "#", "#", "#", "#", "#", ".", "#", "#"},
			[]string{"#", ".", "#", ".", "#", ".", ".", "#", ".", "."},
			[]string{"#", "#", "#", "#", ".", "#", "#", ".", "#", "#"},
			[]string{"#", ".", "#", "#", ".", "#", "#", ".", "#", "#"},
			[]string{"#", ".", "#", "#", "#", "#", "#", ".", "#", "#"},
			[]string{".", ".", "#", ".", "#", ".", ".", ".", ".", "."},
			[]string{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			[]string{"#", ".", "#", "#", "#", "#", "#", "#", ".", "#"},
			[]string{"#", ".", "#", "#", "#", "#", "#", ".", "#", "#"},
		},
		[][]string{
			[]string{"#", ".", "L", "L", ".", "L", "L", ".", "L", "#"},
			[]string{"#", "L", "L", "L", "L", "L", "L", ".", "L", "L"},
			[]string{"L", ".", "L", ".", "L", ".", ".", "L", ".", "."},
			[]string{"L", "L", "L", "L", ".", "L", "L", ".", "L", "L"},
			[]string{"L", ".", "L", "L", ".", "L", "L", ".", "L", "L"},
			[]string{"L", ".", "L", "L", "L", "L", "L", ".", "L", "L"},
			[]string{".", ".", "L", ".", "L", ".", ".", ".", ".", "."},
			[]string{"L", "L", "L", "L", "L", "L", "L", "L", "L", "#"},
			[]string{"#", ".", "L", "L", "L", "L", "L", "L", ".", "L"},
			[]string{"#", ".", "L", "L", "L", "L", "L", ".", "L", "#"},
		},
	}

	for i, start := range stages {
		if i+1 >= len(stages) {
			continue
		}

		next := stages[i+1]

		t.Run(fmt.Sprintf("Stage %v->%v", i, i+1), func(t *testing.T) {
			r := &room{cells: start}

			if got := r.iterate(r.visible, 5).cells; !reflect.DeepEqual(got, next) {
				t.Errorf("iterate() = \n\n%v\n\nwant:\n\n%v", cellsToStr(got), cellsToStr(next))
			}
		})
	}
}

func cellsToStr(cells [][]string) string {
	out := ""

	for _, row := range cells {
		out = out + strings.Join(row, "") + "\n"
	}

	return out
}
