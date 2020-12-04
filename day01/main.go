// --- Day 1: Report Repair ---
//
// After saving Christmas five years in a row, you've decided to take a vacation at a nice
// resort on a tropical island. Surely, Christmas will go on without you.
//
// The tropical island has its own currency and is entirely cash-only. The gold coins used
// there have a little picture of a starfish; the locals just call them stars. None of the
// currency exchanges seem to have heard of them, but somehow, you'll need to find fifty
// of these coins by the time you arrive so you can pay the deposit on your room.
//
// To save your vacation, you need to get all fifty stars by December 25th.
//
// Collect stars by solving puzzles. Two puzzles will be made available on each day in the
// Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle
// grants one star. Good luck!
//
// Before you leave, the Elves in accounting just need you to fix your expense report
// (your puzzle input); apparently, something isn't quite adding up.
//
// Specifically, they need you to find the two entries that sum to 2020 and then multiply
// those two numbers together.
//
// For example, suppose your expense report contained the following:
//
// 1721
// 979
// 366
// 299
// 675
// 1456
//
// In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them
// together produces 1721 * 299 = 514579, so the correct answer is 514579.
//
// Of course, your expense report is much larger. Find the two entries that sum to 2020;
// what do you get if you multiply them together?

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read the input file into a byte array
	cwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/day01/input.txt", cwd))

	if err != nil {
		log.Fatal(err)
	}

	// Convert the input into a string, split into a slice of lines as strings
	input := string(data)
	lines := strings.Split(input, "\n")

	// Convert into a slice of ints
	var nums []int
	for _, line := range lines {
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	// Search for the two numbers adding to 2020
	var match0 int
	var match1 int

	for _, num0 := range nums {
		for _, num1 := range nums {
			sum := num0 + num1

			if sum == 2020 {
				match0 = num0
				match1 = num1
				break
			}
		}
	}

	// If we've found no matches, then these will be the zero value for ints
	if match0 == 0 || match1 == 0 {
		log.Fatal("No matches found")
	}

	// Find the answer: 211899
	answer := match0 * match1
	fmt.Printf("%v", answer)
}
