package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type NumberToken struct {
	line  int
	start int
	end   int
}

type StarToken struct {
	line  int
	start int
}

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

	part2(lines)

	fmt.Printf("Result Part 1: %v \n", partOne(lines))
	fmt.Printf("Result Part 2: %v \n", part2(lines))
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

func part2(lines []string) int {

	var allNumbers []NumberToken
	var allStars []StarToken
	numbers := regexp.MustCompile("\\d+")
	stars := regexp.MustCompile("\\*")

	var num [][]int

	for i, line := range lines {
		num = numbers.FindAllStringIndex(line, -1)
		st := stars.FindAllStringIndex(line, -1)

		for _, n := range num {
			allNumbers = append(allNumbers, NumberToken{line: i, start: n[0], end: n[1]})
		}

		for _, s := range st {
			allStars = append(allStars, StarToken{line: i, start: s[0]})
		}
	}
	fmt.Println(allStars)

	sum := 0
	for _, star := range allStars {
		sum = sum + checkAdjacentNumbers(star.line, star.start, allNumbers, lines)
		println("star: ", star.line, star.start)
	}

	return sum
}

func checkAdjacentNumbers(line int, start int, numbers []NumberToken, lines []string) int {
	var neighbors []string
	for _, number := range numbers {
		if number.line == line || number.line == line-1 || number.line == line+1 {
			if (number.start >= start-3 && number.end > start-1) && (number.start <= start+1) {
				n := ""
				for j := number.start; j < number.end; j++ {
					n += string(lines[number.line][j])
				}

				neighbors = append(neighbors, n)
			}
		}
	}

	if len(neighbors) == 2 {
		println("neighbors: ", neighbors[0], neighbors[1])
		number1, _ := strconv.Atoi(neighbors[0])
		number2, _ := strconv.Atoi(neighbors[1])
		return number1 * number2
	}
	return 0
}
