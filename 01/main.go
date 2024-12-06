package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data, err := os.ReadFile("01/input")
	check(err)

	lines := strings.Split(string(data), "\n")

	left := make([]int, 0, len(lines))
	right := make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		numbers := strings.Split(l, "   ")

		first, err := strconv.Atoi(numbers[0])
		check(err)

		second, err := strconv.Atoi(numbers[1])
		check(err)

		left = append(left, first)
		right = append(right, second)
	}
	slices.Sort(left)
	slices.Sort(right)

	dist := 0
	for i, _ := range left {
		if left[i] > right[i] {
			dist += left[i] - right[i]
		}
		if left[i] < right[i] {
			dist += right[i] - left[i]
		}
	}
	fmt.Printf("The total distance is %v\n", dist)

	similarity := 0
	for i, _ := range left {
		similarity += left[i] * howManyTimes(right, left[i])
	}
	fmt.Printf("The similarity score is %v\n", similarity)
}

func howManyTimes(arraySlice []int, item int) int {
	countr := 0
	for _, arItem := range arraySlice {
		if arItem == item {
			countr += 1
		}
	}
	return countr
}
