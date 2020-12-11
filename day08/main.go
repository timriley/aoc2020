// --- Day 8: Handheld Halting ---
//
// Your flight to the major airline hub reaches cruising altitude without
// incident. While you consider checking the in-flight menu for one of those
// drinks that come with a little umbrella, you are interrupted by the kid
// sitting next to you.
//
// Their handheld game console won't turn on! They ask if you can take a look.
//
// You narrow the problem down to a strange infinite loop in the boot code (your
// puzzle input) of the device. You should be able to fix it, but first you need
// to be able to run the code in isolation.
//
// The boot code is represented as a text file with one instruction per line of
// text. Each instruction consists of an operation (acc, jmp, or nop) and an
// argument (a signed number like +4 or -20).
//
// - acc increases or decreases a single global value called the accumulator by
// the value given in the argument. For example, acc +7 would increase the
// accumulator by 7. The accumulator starts at 0. After an acc instruction, the
// instruction immediately below it is executed next.
// - jmp jumps to a new instruction relative to itself. The next instruction to
// execute is found using the argument as an offset from the jmp instruction;
// for example, jmp +2 would skip the next instruction, jmp +1 would continue to
// the instruction immediately below it, and jmp -20 would cause the instruction
// 20 lines above to be executed next.
// - nop stands for No OPeration - it does nothing. The instruction immediately
// below it is executed next.
//
// For example, consider the following program:
//
// nop +0
// acc +1
// jmp +4
// acc +3
// jmp -3
// acc -99
// acc +1
// jmp -4
// acc +6
//
// These instructions are visited in this order:
//
// nop +0  | 1
// acc +1  | 2, 8(!)
// jmp +4  | 3
// acc +3  | 6
// jmp -3  | 7
// acc -99 |
// acc +1  | 4
// jmp -4  | 5
// acc +6  |
//
// First, the nop +0 does nothing. Then, the accumulator is increased from 0 to
// 1 (acc +1) and jmp +4 sets the next instruction to the other acc +1 near the
// bottom. After it increases the accumulator from 1 to 2, jmp -4 executes,
// setting the next instruction to the only acc +3. It sets the accumulator to
// 5, and jmp -3 causes the program to continue back at the first acc +1.
//
// This is an infinite loop: with this sequence of jumps, the program will run
// forever. The moment the program tries to run any instruction a second time,
// you know it will never terminate.
//
// Immediately before the program would run an instruction a second time, the
// value in the accumulator is 5.
//
// Run your copy of the boot code. Immediately before any instruction is
// executed a second time, what value is in the accumulator?

package main

import (
	"aoc2020/utils/fileinput"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	// line int
	op  string
	arg int
}

func main() {
	var ins []*instruction
	// line := 1
	err := fileinput.LoadThen("day08/input.txt", "\n", func(s string) {
		// ins = append(ins, instructionFromString(line, s))
		ins = append(ins, instructionFromString(s))
		// line++
	})
	if err != nil {
		log.Fatal(err)
	}

	accumulatorBeforeRepeatedInstruction := part1(ins)

	fmt.Printf("Accumulator before repeated instruction (part 1): %v", accumulatorBeforeRepeatedInstruction)
}

func part1(ins []*instruction) int {
	acc := 0
	called := map[int]int{}
	i := 0

	for {
		called[i]++

		if called[i] > 1 {
			return acc
		}

		nextAcc, jump := execute(ins[i], acc)

		acc = nextAcc
		i = i + jump
	}
}

func execute(i *instruction, acc int) (nextAcc int, jump int) {
	switch i.op {
	case "acc":
		return acc + i.arg, 1
	case "jmp":
		return acc, i.arg
	case "nop":
		return acc, 1
	default:
		panic("unknown instruction")
	}
}

// func instructionFromString(line int, s string) *instruction {
func instructionFromString(s string) *instruction {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		log.Fatalf("expected 2 parts in %v", parts)
	}

	arg, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("integer argument expected, received %v", parts[1])
	}

	return &instruction{
		// line: line,
		op:  parts[0],
		arg: arg,
	}
}
