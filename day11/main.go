// --- Day 11: Seating System ---
//
// Your plane lands with plenty of time to spare. The final leg of your journey
// is a ferry that goes directly to the tropical island where you can finally
// start your vacation. As you reach the waiting area to board the ferry, you
// realize you're so early, nobody else has even arrived yet!
//
// By modeling the process people use to choose (or abandon) their seat in the
// waiting area, you're pretty sure you can predict the best place to sit. You
// make a quick map of the seat layout (your puzzle input).
//
// The seat layout fits neatly on a grid. Each position is either floor (.), an
// empty seat (L), or an occupied seat (#). For example, the initial seat layout
// might look like this:
//
// L.LL.LL.LL
// LLLLLLL.LL
// L.L.L..L..
// LLLL.LL.LL
// L.LL.LL.LL
// L.LLLLL.LL
// ..L.L.....
// LLLLLLLLLL
// L.LLLLLL.L
// L.LLLLL.LL
//
// Now, you just need to model the people who will be arriving shortly.
// Fortunately, people are entirely predictable and always follow a simple set
// of rules. All decisions are based on the number of occupied seats adjacent to
// a given seat (one of the eight positions immediately up, down, left, right,
// or diagonal from the seat). The following rules are applied to every seat
// simultaneously:
//
// - If a seat is empty (L) and there are no occupied seats adjacent to it, the
// seat becomes occupied.
// - If a seat is occupied (#) and four or more seats adjacent to it are also
// occupied, the seat becomes empty.
// - Otherwise, the seat's state does not change.
// - Floor (.) never changes; seats don't move, and nobody sits on the floor.
//
// After one round of these rules, every seat in the example layout becomes occupied:
//
// #.##.##.##
// #######.##
// #.#.#..#..
// ####.##.##
// #.##.##.##
// #.#####.##
// ..#.#.....
// ##########
// #.######.#
// #.#####.##
//
// After a second round, the seats with four or more occupied adjacent seats
// become empty again:
//
// #.LL.L#.##
// #LLLLLL.L#
// L.L.L..L..
// #LLL.LL.L#
// #.LL.LL.LL
// #.LLLL#.##
// ..L.L.....
// #LLLLLLLL#
// #.LLLLLL.L
// #.#LLLL.##
//
// This process continues for three more rounds:
//
// #.##.L#.##
// #L###LL.L#
// L.#.#..#..
// #L##.##.L#
// #.##.LL.LL
// #.###L#.##
// ..#.#.....
// #L######L#
// #.LL###L.L
// #.#L###.##

// #.#L.L#.##
// #LLL#LL.L#
// L.L.L..#..
// #LLL.##.L#
// #.LL.LL.LL
// #.LL#L#.##
// ..L.L.....
// #L#LLLL#L#
// #.LLLLLL.L
// #.#L#L#.##

// #.#L.L#.##
// #LLL#LL.L#
// L.#.L..#..
// #L##.##.L#
// #.#L.LL.LL
// #.#L#L#.##
// ..L.L.....
// #L#L##L#L#
// #.LLLLLL.L
// #.#L#L#.##
//
// At this point, something interesting happens: the chaos stabilizes and
// further applications of these rules cause no seats to change state! Once
// people stop moving around, you count 37 occupied seats.
//
// Simulate your seating area by applying the seating rules repeatedly until no
// seats change state. How many seats end up occupied?

package main

import (
	"aoc2020/utils/fileinput"
	"fmt"
	"strings"
)

type room struct {
	cells [][]string
}

func main() {
	var positions [][]string
	err := fileinput.LoadThen("day11/input.txt", "\n", func(s string) {
		var row []string
		for _, rowStr := range strings.Split(s, "\n") {
			for _, c := range rowStr {
				row = append(row, string(c))
			}
		}
		positions = append(positions, row)
	})
	if err != nil {
		panic(err)
	}

	r := newRoom(positions)

	res := part1(r)

	fmt.Println(res) // 2283
}

func part1(r *room) int {
	occupiedCount := r.occupiedCount()

	for {
		r = r.iterate()

		if occupiedCount == r.occupiedCount() {
			return occupiedCount
		}

		occupiedCount = r.occupiedCount()
	}
}

func newRoom(cells [][]string) *room {
	return &room{cells}
}

func (r *room) iterate() *room {
	newCells := [][]string{}
	for _, row := range r.cells {
		newCells = append(newCells, make([]string, len(row)))
	}

	for x, row := range r.cells {
		for y, cell := range row {
			occupied := 0
			for _, n := range r.neighbours(x, y) {
				if n == "#" {
					occupied++
				}
			}

			newCell := cell

			// If a seat is empty (L) and there are no occupied seats adjacent to it, the
			// seat becomes occupied
			if cell == "L" && occupied == 0 {
				newCell = "#"
			}

			// If a seat is occupied (#) and four or more seats adjacent to it are also
			// occupied, the seat becomes empty
			if cell == "#" && occupied >= 4 {
				newCell = "L"
			}

			newCells[x][y] = newCell
		}
	}

	return newRoom(newCells)
}

func (r *room) occupiedCount() int {
	count := 0
	for _, row := range r.cells {
		for _, cell := range row {
			if cell == "#" {
				count++
			}
		}
	}
	return count
}

func (r *room) neighbours(x, y int) []string {
	checks := map[int][]int{
		x - 1: {y - 1, y, y + 1},
		x:     {y - 1, y + 1},
		x + 1: {y - 1, y, y + 1},
	}

	neighbours := []string{}

	for xCheck, yChecks := range checks {
		if xCheck >= 0 && xCheck < len(r.cells) {
			row := r.cells[xCheck]

			for _, yCheck := range yChecks {
				if yCheck >= 0 && yCheck < len(row) {
					neighbours = append(neighbours, row[yCheck])
				}
			}
		}
	}

	return neighbours
}
