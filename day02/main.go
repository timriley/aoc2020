//--- Day 2: Password Philosophy ---
//
// Your flight departs in a few days from the coastal airport; the easiest way down to the coast from here is via
// toboggan.
//
// The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day. "Something's wrong with our computers; we
// can't log in!" You ask if you can take a look.
//
// Their password database seems to be a little corrupted: some of the passwords wouldn't have been allowed by the
// Official Toboggan Corporate Policy that was in effect when they were chosen.
//
// To try to debug the problem, they have created a list (your puzzle input) of passwords (according to the corrupted
// database) and the corporate policy when that password was set.
//
// For example, suppose you have the following list:
//
// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc
//
// Each line gives the password policy and then the password. The password policy indicates the lowest and highest
// number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password
// must contain a at least 1 time and at most 3 times.
//
// In the above example, 2 passwords are valid. The middle password, cdefg, is not; it contains no instances of b, but
// needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of
// their respective policies.
//
// How many passwords are valid according to their policies?
//
// --- Part Two ---
//
// While it appears you validated the passwords correctly, they don't seem to be what the Official Toboggan Corporate
// Authentication System is expecting.
//
// The shopkeeper suddenly realizes that he just accidentally explained the password policy rules from his old job at
// the sled rental place down the street! The Official Toboggan Corporate Policy actually works a little differently.
//
// Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second
// character, and so on. (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these
// positions must contain the given letter. Other occurrences of the letter are irrelevant for the purposes of policy
// enforcement.
//
// Given the same example list from above:
//
// 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
// 1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
// 2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
//
// How many passwords are valid according to the new interpretation of the policies?

package main

import (
	"aoc2020/utils/str"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passwordEntry struct {
	character string
	num0      int
	num1      int
	password  string
}

// Example entry: "8-11 m: wzxcmwgmmvmgq"
var entryLineRegexp = regexp.MustCompile(`^(?P<minTimes>\d+)-(?P<maxTimes>\d+)\s(?P<character>[a-z]):\s(?P<password>[a-z]+)$`)

func main() {
	entries, _ := parseInput("day02/input.txt")

	part1ValidPasswordCount := part1(entries)
	part2ValidPasswordCount := part2(entries)

	fmt.Printf("Part 1 valid password count: %v\n", part1ValidPasswordCount)
	fmt.Printf("Part 2 valid password count: %v\n", part2ValidPasswordCount)
}

func part1(entries []passwordEntry) int {
	count := 0

	for _, e := range entries {
		if e.passwordContainsCharacterEnoughTimes() {
			count++
		}
	}

	return count
}

func part2(entries []passwordEntry) int {
	count := 0

	for _, e := range entries {
		if e.passwordHasCharacterInSingularPosition() {
			count++
		}
	}

	return count
}

func parseInput(fileName string) (entries []passwordEntry, err error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return entries, err
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts, ok := parseEntry(line)

		if !ok {
			continue
		}

		minTimes, _ := strconv.Atoi(parts["minTimes"])
		maxTimes, _ := strconv.Atoi(parts["maxTimes"])

		entry := passwordEntry{
			character: parts["character"],
			num0:      minTimes,
			num1:      maxTimes,
			password:  parts["password"],
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

// Thanks, https://stackoverflow.com/a/53587770/308563
func parseEntry(line string) (map[string]string, bool) {
	result := make(map[string]string)

	matchCheck := entryLineRegexp.MatchString(line)

	if !matchCheck {
		return result, false
	}

	match := entryLineRegexp.FindStringSubmatch(line)

	for i, val := range match {
		key := entryLineRegexp.SubexpNames()[i]
		result[key] = val
	}

	return result, true
}

func (e passwordEntry) passwordContainsCharacterEnoughTimes() bool {
	count := strings.Count(e.password, e.character)
	return count >= e.num0 && count <= e.num1
}

func (e passwordEntry) passwordHasCharacterInSingularPosition() bool {
	char0, err0 := str.CharAt(e.password, e.num0-1)
	char1, err1 := str.CharAt(e.password, e.num1-1)

	return err0 == nil && err1 == nil && (char0 == e.character) != (char1 == e.character)
}
