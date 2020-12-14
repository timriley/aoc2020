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
//
// --- Part Two ---
//
// As soon as people start to arrive, you realize your mistake. People don't
// just care about adjacent seats - they care about the first seat they can see
// in each of those eight directions!
//
// Now, instead of considering just the eight immediately adjacent seats,
// consider the first seat in each of those eight directions. For example, the
// empty seat below would see eight occupied seats:
//
// .......#.
// ...#.....
// .#.......
// .........
// ..#L....#
// ....#....
// .........
// #........
// ...#.....
//
// The leftmost empty seat below would only see one empty seat, but cannot see
// any of the occupied ones:
//
// .............
// .L.L.#.#.#.#.
// .............
//
// The empty seat below would see no occupied seats:
//
// .##.##.
// #.#.#.#
// ##...##
// ...L...
// ##...##
// #.#.#.#
// .##.##.
//
// Also, people seem to be more tolerant than you expected: it now takes five or
// more visible occupied seats for an occupied seat to become empty (rather than
// four or more from the previous rules). The other rules still apply: empty
// seats that see no occupied seats become occupied, seats matching no rule
// don't change, and floor never changes.
//
// Given the same starting layout as above, these new rules cause the seating
// area to shift around as follows:
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
// #.LL.LL.L#
// #LLLLLL.LL
// L.L.L..L..
// LLLL.LL.LL
// L.LL.LL.LL
// L.LLLLL.LL
// ..L.L.....
// LLLLLLLLL#
// #.LLLLLL.L
// #.LLLLL.L#
//
// #.L#.##.L#
// #L#####.LL
// L.#.#..#..
// ##L#.##.##
// #.##.#L.##
// #.#####.#L
// ..#.#.....
// LLL####LL#
// #.L#####.L
// #.L####.L#
//
// #.L#.L#.L#
// #LLLLLL.LL
// L.L.L..#..
// ##LL.LL.L#
// L.LL.LL.L#
// #.LLLLL.LL
// ..L.L.....
// LLLLLLLLL#
// #.LLLLL#.L
// #.L#LL#.L#
//
// #.L#.L#.L#
// #LLLLLL.LL
// L.L.L..#..
// ##L#.#L.L#
// L.L#.#L.L#
// #.L####.LL
// ..#.#.....
// LLL###LLL#
// #.LLLLL#.L
// #.L#LL#.L#
//
// #.L#.L#.L#
// #LLLLLL.LL
// L.L.L..#..
// ##L#.#L.L#
// L.L#.LL.L#
// #.LLLL#.LL
// ..#.L.....
// LLL###LLL#
// #.LLLLL#.L
// #.L#LL#.L#
//
// Again, at this point, people stop shifting around and the seating area
// reaches equilibrium. Once this occurs, you count 26 occupied seats.
//
// Given the new visibility method and the rule change for occupied seats
// becoming empty, once equilibrium is reached, how many seats end up occupied?

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

	part1Visible := part1(r) // 2283
	part2Visible := part2(r) // 2054

	fmt.Printf("Visible (part 1): %v\n", part1Visible)
	fmt.Printf("Visible (part 2): %v\n", part2Visible)
}

func part1(r *room) int {
	occupiedCount := r.occupiedCount()

	for {
		r = r.iterate(r.neighbours, 4)

		if occupiedCount == r.occupiedCount() {
			return occupiedCount
		}

		occupiedCount = r.occupiedCount()
	}
}

func part2(r *room) int {
	occupiedCount := r.occupiedCount()

	for {
		r = r.iterate(r.visible, 5)

		if occupiedCount == r.occupiedCount() {
			return occupiedCount
		}

		occupiedCount = r.occupiedCount()
	}
}

func newRoom(cells [][]string) *room {
	return &room{cells}
}

func (r *room) iterate(finder func(x, y int) []string, occupiedMin int) *room {
	newCells := [][]string{}
	for _, row := range r.cells {
		newCells = append(newCells, make([]string, len(row)))
	}

	for x, row := range r.cells {
		for y, cell := range row {
			occupied := 0
			for _, n := range finder(x, y) {
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

			// If a seat is occupied (#) and [occupidMin] or more seats adjacent to it are
			// also occupied, the seat becomes empty
			if cell == "#" && occupied >= occupiedMin {
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

func (r *room) visible(x, y int) []string {
	finders := []func(x, y int) (int, int){
		// straight right
		func(x, y int) (int, int) {
			return x, y + 1
		},
		// diagonal down and to the right
		func(x, y int) (int, int) {
			return x + 1, y + 1
		},
		// straight down
		func(x, y int) (int, int) {
			return x + 1, y
		},
		// diagonal down and to the left
		func(x, y int) (int, int) {
			return x + 1, y - 1
		},
		// straight left
		func(x, y int) (int, int) {
			return x, y - 1
		},
		// diagonal up and to the left
		func(x, y int) (int, int) {
			return x - 1, y - 1
		},
		// straight up
		func(x, y int) (int, int) {
			return x - 1, y
		},
		// diagonal up and to the right
		func(x, y int) (int, int) {
			return x - 1, y + 1
		},
	}

	visible := []string{}
	for _, f := range finders {
		findX, findY := x, y
		for {
			findX, findY = f(findX, findY)

			if findX < 0 || findX >= len(r.cells) {
				break
			}

			row := r.cells[findX]

			if findY < 0 || findY >= len(row) {
				break
			}

			cell := row[findY]

			if cell == "." {
				continue
			}

			visible = append(visible, cell)
			break
		}
	}
	return visible
}
