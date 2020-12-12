package main

import (
	"reflect"
	"testing"
)

func TestContiguousAddendsForSum(t *testing.T) {
	nums := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	targetSum := 127

	want := []int{15, 25, 47, 40}
	result := contiguousAddendsForSum(nums, targetSum)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("contiguousAddendsForSum(...) = %v; want %v", result, want)
	}
}
