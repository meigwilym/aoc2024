package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput(filename string) string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := os.ReadFile(pwd + "/" + filename)
	check(err)

	return string(data)
}

func main() {
	text := getInput("input")

	total := getTotalMultiplied(text)

	fmt.Println("Total multiplied", total)
	fmt.Println("Total processed", getTotalProcessed(text))
}

func getTotalMultiplied(text string) int {
	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	allMatches := r.FindAllStringSubmatch(text, -1)

	total := 0
	for _, match := range allMatches {
		val1, err := strconv.Atoi(match[1])
		check(err)
		val2, err := strconv.Atoi(match[2])
		check(err)
		total += val1 * val2
	}
	return total
}

func getTotalProcessed(text string) int {

	r, _ := regexp.Compile("(mul\\((\\d{1,3}),(\\d{1,3})\\)|(don't\\(\\))|do\\(\\))")

	allMatches := r.FindAllStringSubmatch(text, -1)
	capture := true
	total := 0

	for _, match := range allMatches {
		if match[4] == "don't()" {
			capture = false
		}
		if match[1] == "do()" {
			capture = true
		}
		if capture == true {
			if match[2] != "" {
				val1, err := strconv.Atoi(match[2])
				check(err)
				val2, err := strconv.Atoi(match[3])
				check(err)
				total += val1 * val2
			}
		}
	}
	return total
}
