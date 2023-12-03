package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("Result Part 1: %v \n", partOne(lines))
	// fmt.Printf("Result Part 2: %v \n", partTwo(lines))
}

func partOne(lines []string) int {
	length := len(lines[0])
	sum := 0

	for i, line := range lines {
		numberStart := -1
		numberEnd := -1
		digit := ""

		for j, char := range line {
			if unicode.IsDigit(char) {
				if numberStart == -1 {
					numberStart = j
					numberEnd = j
				}

				if !(j+1 > length-1) && unicode.IsDigit(rune(line[j+1])) {
					numberEnd = j + 1
				} else {
					valid := checkSurroundings(lines, i, numberStart, numberEnd)

					if valid {
						for i := numberStart; i <= numberEnd; i++ {
							digit = digit + string(line[i])
						}

						if digit != "" {
							partNumber, err := strconv.Atoi(digit)
							if err != nil {
								log.Fatal("ERROR")
							}

							sum = sum + partNumber
						}
					} else {
						for i := numberStart; i <= numberEnd; i++ {
							digit = digit + string(line[i])
						}
						fmt.Printf("%v \n", digit)
					}
					numberStart = -1
					numberEnd = -1
					digit = ""
				}
			}
		}
	}

	return sum
}

func checkSurroundings(lines []string, line int, numberStart int, numberEnd int) bool {

	if !(numberStart-1 < 0) && !unicode.IsDigit(rune(lines[line][numberStart-1])) && string(lines[line][numberStart-1]) != "." {
		return true
	}
	if !(numberEnd+1 > len(lines[line])-1) && !unicode.IsDigit(rune(lines[line][numberEnd+1])) && string(lines[line][numberEnd+1]) != "." {
		return true
	}
	if !(numberStart-1 < 0) && !(line-1 < 0) && !unicode.IsDigit(rune(lines[line-1][numberStart-1])) && string(lines[line-1][numberStart-1]) != "." {
		return true
	}
	if !(numberEnd+1 > len(lines[line])-1) && !(line-1 < 0) && !unicode.IsDigit(rune(lines[line-1][numberEnd+1])) && string(lines[line-1][numberEnd+1]) != "." {
		return true
	}
	if !(numberStart-1 < 0) && !(line+1 > len(lines)-1) && !unicode.IsDigit(rune(lines[line+1][numberStart-1])) && string(lines[line+1][numberStart-1]) != "." {
		return true
	}
	if !(numberEnd+1 > len(lines[line])-1) && !(line+1 > len(lines)-1) && !unicode.IsDigit(rune(lines[line+1][numberEnd+1])) && string(lines[line+1][numberEnd+1]) != "." {
		return true
	}

	for i := numberStart; i <= numberEnd; i++ {
		if !(line+1 > len(lines)-1) && !unicode.IsDigit(rune(lines[line+1][i])) && string(lines[line+1][i]) != "." {
			return true
		}
		if !(line-1 < 0) && !unicode.IsDigit(rune(lines[line-1][i])) && string(lines[line-1][i]) != "." {
			return true
		}
	}

	return false
}
