//--- Day 2: Password Philosophy ---
//
// Your flight departs in a few days from the coastal airport; the easiest way down to the
// coast from here is via toboggan.
//
// The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day. "Something's
// wrong with our computers; we can't log in!" You ask if you can take a look.
//
// Their password database seems to be a little corrupted: some of the passwords wouldn't
// have been allowed by the Official Toboggan Corporate Policy that was in effect when
// they were chosen.
//
// To try to debug the problem, they have created a list (your puzzle input) of passwords
// (according to the corrupted database) and the corporate policy when that password was
// set.
//
// For example, suppose you have the following list:
//
// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc
//
// Each line gives the password policy and then the password. The password policy
// indicates the lowest and highest number of times a given letter must appear for the
// password to be valid. For example, 1-3 a means that the password must contain a at
// least 1 time and at most 3 times.
//
// In the above example, 2 passwords are valid. The middle password, cdefg, is not; it
// contains no instances of b, but needs at least 1. The first and third passwords are
// valid: they contain one a or nine c, both within the limits of their respective
// policies.
//
// How many passwords are valid according to their policies?

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passwordEntry struct {
	character string
	minTimes  int
	maxTimes  int
	password  string
}

// Example entry: "8-11 m: wzxcmwgmmvmgq"
var entryLineRegexp = regexp.MustCompile(`^(?P<minTimes>\d+)-(?P<maxTimes>\d+)\s(?P<character>[a-z]):\s(?P<password>[a-z]+)$`)

func main() {
	entries, _ := parseInput("day02/input.txt")

	validPasswordCount := part1(entries)

	fmt.Printf("Valid password count: %v", validPasswordCount)
}

func part1(entries []passwordEntry) int {
	count := 0

	for _, e := range entries {
		if e.isPasswordValid() {
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
			minTimes:  minTimes,
			maxTimes:  maxTimes,
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

func (e passwordEntry) isPasswordValid() bool {
	count := 0

	for _, c := range e.password {
		if string(c) == e.character {
			count++
		}
	}

	return count >= e.minTimes && count <= e.maxTimes
}
