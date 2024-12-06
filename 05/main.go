package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := "input"
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "test" {
			inputFile = "test"
		}
	}
	lines := getInput(inputFile)

	updateFlag := false
	rules := [][2]int{}
	updates := [][]int{}
	for _, line := range lines {
		if line == "" {
			updateFlag = true
			continue
		}

		if updateFlag {
			updateValues := strings.Split(line, ",")
			updateInts := []int{}
			for _, val := range updateValues {
				updateInt, _ := strconv.Atoi(val)
				updateInts = append(updateInts, updateInt)
			}
			updates = append(updates, updateInts)
		} else {
			rulesValues := strings.Split(line, "|")
			firstRule, _ := strconv.Atoi(rulesValues[0])
			secondRule, _ := strconv.Atoi(rulesValues[1])
			theseRules := [2]int{firstRule, secondRule}
			rules = append(rules, theseRules)
		}
	}

	fmt.Println(rules)
	fmt.Println(updates)

}

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

	data, err := os.ReadFile(pwd + "/" + filename)
	check(err)

	return strings.Split(string(data), "\n")
}
