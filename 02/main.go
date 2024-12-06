package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput(filename string) []string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	data, err := os.ReadFile(pwd + "/" + filename)
	check(err)

	return strings.Split(string(data), "\n")
}

func main() {
	lines := getInput("input")

	safeReports := 0

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		reports := make([]int, 0)
		reportsStr := strings.Split(l, " ")

		for _, report := range reportsStr {
			val, err := strconv.Atoi(report)
			check(err)
			reports = append(reports, val)
		}
		levelReports := getReportsWithOneLevelMissing(reports)

		isSafe := false
		if isIncreasing(reports) || isDecreasing(reports) {
			if diffLeastOne(reports) && diffMostThree(reports) {
				isSafe = true
				safeReports += 1
			}
		}
		if isSafe == false {
			for _, report := range levelReports {
				if isIncreasing(report) || isDecreasing(report) {
					if diffLeastOne(report) && diffMostThree(report) {
						safeReports += 1
						break
					}
				}
			}
		}
	}
	fmt.Printf("There are %v safe reports\n", safeReports)

}

func isIncreasing(reports []int) bool {
	for i, report := range reports {
		if i > 0 {
			if report < reports[i-1] {
				return false
			}
		}
	}
	return true
}

func isDecreasing(reports []int) bool {
	for i, report := range reports {
		if i > 0 {
			if report > reports[i-1] {
				return false
			}
		}
	}
	return true
}

func diffLeastOne(reports []int) bool {
	for i, _ := range reports {
		if i == 0 {
			continue
		}
		if reports[i-1]-reports[i] == 0 {
			return false
		}
	}
	return true
}

func diffMostThree(reports []int) bool {
	for i, _ := range reports {
		if i == 0 {
			continue
		}
		diff := reports[i-1] - reports[i]
		if diff > 3 || diff < -3 {
			return false
		}
	}
	return true
}

func getReportsWithOneLevelMissing(report []int) [][]int {
	returnedArray := make([][]int, len(report))
	for index, item := range report {
		for i := 0; i < len(report); i++ {
			if i != index {
				returnedArray[i] = append(returnedArray[i], item)
			}
		}
	}
	return returnedArray
}
