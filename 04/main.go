package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	xmas string = "XMAS"
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

	data, err := os.ReadFile(pwd + "/" + filename)
	check(err)

	return strings.Split(string(data), "\n")
}

func main() {
	lines := getInput("input")

	instances := countXmas(lines)

	fmt.Printf("There are %d instances of xmas\n", instances)

	patternInst := countPattern(lines)

	fmt.Printf("There are %d instances of the pattern\n", patternInst)

}

func countPattern(lines []string) int {
	lineLength := utf8.RuneCountInString(lines[0])
	counter := 0
	for i, line := range lines {
		if line == "" {
			continue
		}
		for j, char := range strings.Split(line, "") {
			if i > 0 && i < len(lines)-2 && j > 0 && j < lineLength-1 {
				if char != "A" {
					continue
				}
				if (string(lines[i-1][j-1])+char+string(lines[i+1][j+1]) == "MAS" || string(lines[i-1][j-1])+char+string(lines[i+1][j+1]) == "SAM") &&
					(string(lines[i-1][j+1])+char+string(lines[i+1][j-1]) == "MAS" || string(lines[i-1][j+1])+char+string(lines[i+1][j-1]) == "SAM") {
					counter++
				}
			}
		}
	}
	return counter
}

func countXmas(lines []string) int {
	lineLength := utf8.RuneCountInString(lines[0])
	counter := 0
	for i, line := range lines {
		if line == "" {
			continue
		}
		for j, char := range strings.Split(line, "") {

			// right
			if j < lineLength-3 && string(line[j:j+4]) == xmas {
				counter++
			}
			// left
			if j > 2 && char+string(line[j-1])+string(line[j-2])+string(line[j-3]) == xmas {
				counter++
			}
			// up
			if i > 2 && char+string(lines[i-1][j])+string(lines[i-2][j])+string(lines[i-3][j]) == xmas {
				counter++
			}
			// down
			if i < len(lines)-4 && char+string(lines[i+1][j])+string(lines[i+2][j])+string(lines[i+3][j]) == xmas {
				counter++
			}

			// down right
			if i < len(lines)-4 && j < lineLength-3 {
				if char+string(lines[i+1][j+1])+string(lines[i+2][j+2])+string(lines[i+3][j+3]) == xmas {
					counter++
				}
			}
			// down left
			if i < len(lines)-4 && j > 2 {
				if char+string(lines[i+1][j-1])+string(lines[i+2][j-2])+string(lines[i+3][j-3]) == xmas {
					counter++
				}

			}
			// up left
			if i > 2 && j > 2 {
				if char+string(lines[i-1][j-1])+string(lines[i-2][j-2])+string(lines[i-3][j-3]) == xmas {
					counter++
				}

			}
			// up right
			if i > 2 && j < lineLength-3 {
				if char+string(lines[i-1][j+1])+string(lines[i-2][j+2])+string(lines[i-3][j+3]) == xmas {
					counter++
				}
			}
		}
	}
	return counter
}
