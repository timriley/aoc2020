// --- Day 3: Toboggan Trajectory ---
//
// With the toboggan login problems resolved, you set off toward the airport. While travel by toboggan might be easy,
// it's certainly not safe: there's very minimal steering and the area is covered in trees. You'll need to see which
// angles will take you near the fewest trees.
//
// Due to the local geology, trees in this area only grow on exact integer coordinates in a grid. You make a map (your
// puzzle input) of the open squares (.) and trees (#) you can see. For example:
//
// ..##.......
// #...#...#..
// .#....#..#.
// ..#.#...#.#
// .#...##..#.
// ..#.##.....
// .#.#.#....#
// .#........#
// #.##...#...
// #...##....#
// .#..#...#.#
//
// These aren't the only trees, though; due to something you read about once involving arboreal genetics and biome
// stability, the same pattern repeats to the right many times:
//
// ..##.........##.........##.........##.........##.........##.......  --->
// #...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
// .#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
// ..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
// .#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
// ..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....  --->
// .#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
// .#........#.#........#.#........#.#........#.#........#.#........#
// #.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
// #...##....##...##....##...##....##...##....##...##....##...##....#
// .#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
//
// You start on the open square (.) in the top-left corner and need to reach the bottom (below the bottom-most row on
// your map).
//
// The toboggan can only follow a few specific slopes (you opted for a cheaper model that prefers rational numbers);
// start by counting all the trees you would encounter for the slope right 3, down 1:
//
// From your starting position at the top-left, check the position that is right 3 and down 1. Then, check the position
// that is right 3 and down 1 from there, and so on until you go past the bottom of the map.
//
// The locations you'd check in the above example are marked here with O where there was an open square and X where
// there was a tree:
//
// ..##.........##.........##.........##.........##.........##.......  --->
// #..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
// .#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
// ..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
// .#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
// ..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->
// .#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
// .#........#.#........X.#........#.#........#.#........#.#........#
// #.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...
// #...##....##...##....##...#X....##...##....##...##....##...##....#
// .#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
//
// In this example, traversing the map using this slope would cause you to encounter 7 trees.
//
// Starting at the top-left corner of your map and following a slope of right 3 and down 1, how many trees would you
// encounter?
//
// --- Part Two ---
//
// Time to check the rest of the slopes - you need to minimize the probability of a sudden arboreal stop, after all.
//
// Determine the number of trees you would encounter if, for each of the following slopes, you start at the top-left
// corner and traverse the map all the way to the bottom:
//
// Right 1, down 1.
// Right 3, down 1. (This is the slope you already checked.)
// Right 5, down 1.
// Right 7, down 1.
// Right 1, down 2.
//
// In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s) respectively; multiplied together, these
// produce the answer 336.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Row []string

func main() {
	rows, err := parseInput("day03/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	part1Trees := part1(rows)
	part2Trees := part2(rows)

	fmt.Printf("Part 1: %v trees\nt", part1Trees)
	fmt.Printf("Part 2: %v trees\nt", part2Trees)
}

func part1(rows []Row) int {
	return treesForTraversal(rows, 1, 3)
}

func part2(rows []Row) int {
	type TraverseBy struct {
		downBy  int
		rightBy int
	}

	traversals := []TraverseBy{
		{
			downBy:  1,
			rightBy: 1,
		},
		{
			downBy:  1,
			rightBy: 3,
		},
		{
			downBy:  1,
			rightBy: 5,
		},
		{
			downBy:  1,
			rightBy: 7,
		},
		{
			downBy:  2,
			rightBy: 1,
		},
	}

	m := 1

	for _, t := range traversals {
		m = m * treesForTraversal(rows, t.downBy, t.rightBy)
	}

	return m
}

func treesForTraversal(rows []Row, downBy int, rightBy int) int {
	treesCount := 0

	for x, y := 0, 0; x < len(rows); x, y = x+downBy, y+rightBy {
		cell := rows[x].CellAt(y)

		if cell == "#" {
			treesCount++
		}
	}

	return treesCount
}

func (r Row) CellAt(index int) string {
	if index < len(r) {
		return r[index]
	}

	wrappedIndex := index % len(r)
	return r[wrappedIndex]
}

func parseInput(fileName string) ([]Row, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return make([]Row, 0), err
	}

	lines := strings.Split(string(data), "\n")

	var rows []Row

	for _, line := range lines {
		if line == "" {
			continue
		}

		var row []string

		for _, r := range []rune(line) {
			row = append(row, string(r))
		}

		rows = append(rows, row)
	}

	return rows, nil
}
